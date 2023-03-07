#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>

#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 9999

int main() {
    int sockfd;
    struct sockaddr_in serv_addr;
    char buffer[1024] = {0};
    
    // Create socket file descriptor
    if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("socket failed");
        exit(EXIT_FAILURE);
    }
    
    // Set server address and port
    serv_addr.sin_family = AF_INET;
    serv_addr.sin_port = htons(SERVER_PORT);
    if (inet_pton(AF_INET, SERVER_IP, &serv_addr.sin_addr) <= 0) {
        perror("inet_pton failed");
        exit(EXIT_FAILURE);
    }
    
    // Connect to server
    if (connect(sockfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) {
        perror("connect failed");
        exit(EXIT_FAILURE);
    }
    
    // Receive output from server
    while (1) {
        ssize_t num_bytes = recv(sockfd, buffer, sizeof(buffer), 0);
        if (num_bytes < 0) {
            perror("recv failed");
            close(sockfd);
            exit(EXIT_FAILURE);
        }
        if (num_bytes == 0) {
            printf("Connection closed by server.\n");
            close(sockfd);
            break;
        }
        printf("%s", buffer);
        fflush(stdout);
    }
    
    // Reconnect to server
    while (1) {
        printf("Reconnecting to server...\n");
        sleep(1);
        if ((sockfd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
            perror("socket failed");
            continue;
        }
        if (connect(sockfd, (struct sockaddr *)&serv_addr, sizeof(serv_addr)) < 0) {
            perror("connect failed");
            close(sockfd);
            continue;
        }
        printf("Connected to server.\n");
        // Receive output from server (same as above)
        // ...
        break;
    }
    
    return 0;
}