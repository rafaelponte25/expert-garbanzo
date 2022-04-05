#!/bin/bash

# Use \n in printf to skip a line for the cursor
printf "Initializing memo...\n"

# -n will not put the cursor on the next line
echo -n "What is your name? "
read USER_INPUT

function welcome_message {
    echo "Hello ${USER_INPUT}, I am Dexter"
}

next_message() {
    echo "I am preparing you for something..."
}

# Run the functions only after the are declared
welcome_message
next_message