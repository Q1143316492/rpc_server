#include <bits/stdc++.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <assert.h>
#include <unistd.h>

#pragma pack(1)
struct Pack {
    int32_t dataLen;
    int32_t dataType;
    char*   data;
};
#pragma pack()

int judge() {
    union {
        short value;
        char union_bytes[sizeof( short ) ];
    } test;
    test.value = 0x0102;
    printf("%d %d %d\n", sizeof(short) ,test.union_bytes[0], test.union_bytes[1]);
    if(test.union_bytes[0] == 1 && test.union_bytes[1] == 2) {
        puts("大端序");
        return 1;
    } else if(test.union_bytes[0] == 2 && test.union_bytes[1] == 1) {
        puts("小端序");
        return 2;
    } else {
        puts("未知");
        return 0;
    }
}

int main(int argc, char *argv[])
{
    if(argc <= 2)
    {
        printf("usage: %s ip_address port_number\n", basename( argv[0] ));
        // basename :
        // Return the file name within directory of FILENAME.  We don't
        //  declare the function if the `basename' macro is available (defined
        //  in <libgen.h>) which makes the XPG version of this function
        //  available.
        return 1;
    }

    const char* ip = argv[1];
    int port = atoi( argv[2] );

    printf("ip = %s port = %d\n", ip, port);

    struct sockaddr_in server_address;
    bzero(&server_address, sizeof( server_address ));
    server_address.sin_family = AF_INET;    //#define PF_INET 2 //IP protocol family. // PF_INET equal AF_INET

    inet_pton( AF_INET, ip, &server_address.sin_addr );
    server_address.sin_port = htons( port );

    int sockfd = socket( PF_INET, SOCK_STREAM, 0 );
    assert( sockfd >= 0 );

    int ret = judge();

    if( connect( sockfd, ( struct sockaddr* )&server_address, sizeof( server_address ) ) < 0 )
    {
        printf("connection failed\n");
    }
    else
    {
        char buf[1024] = "hello";

        int32_t dataLen  = htonl(5);
        int32_t dataType = htonl(1001);
        write(sockfd, &dataLen,  4);
        write(sockfd, &dataType, 4);
        write(sockfd, buf, 5);
        
        // write(sockfd, buf, 11);
    }
    close(sockfd);
    return 0;
}
