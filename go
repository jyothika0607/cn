#include<stdio.h>
 
int main()
{
    int w, i, f, frames[50];
 
    printf("Enter window size: ");
    scanf("%d", &w);
 
    printf("\nEnter number of frames to transmit: ");
    scanf("%d", &f); 
 
    printf("\nEnter %d frames: ", f);
 
    for (i = 1; i <= f; i++)
        scanf("%d", &frames[i]);
 
    printf("\nWith sliding window protocol the frames will be sent in the following manner (assuming no corruption of frames)\n\n");
    printf("After sending %d frames at each stage sender waits for acknowledgement sent by the receiver\n\n", w);
 
    for (i = 1; i <= f; i++) {
        if (i % w == 0) {
            printf("%d\n", frames[i]);
            printf("Acknowledgement of above frames sent is received by sender\n\n");
        } else {
            printf("%d ", frames[i]);
        }
    }
 
    if (f % w != 0)
        printf("\nAcknowledgement of above frames sent is received by sender\n");
 
    return 0;
}

/*
Example Execution:

Input:
Enter window size: 3

Enter number of frames to transmit: 7

Enter 7 frames: 1 2 3 4 5 6 7

Output:
With sliding window protocol the frames will be sent in the following manner (assuming no corruption of frames)

After sending 3 frames at each stage sender waits for acknowledgement sent by the receiver

1 2 3
Acknowledgement of above frames sent is received by sender

4 5 6
Acknowledgement of above frames sent is received by sender

7
Acknowledgement of above frames sent is received by sender
*/
