## Table of Contents

- [Introduction](#introduction)
- [How to change ROMs](#how-to-change-roms)
- [Test ROMs](#test-roms)
- [Help Code](#help-code)
- [Problems](#problems)

## Introduction
This project is meant to create a Chip-8 Emulator where you can load ROMs and play games like Pong or many others.

## How to change ROMs
You can load a ROM in **game/game.go** → **RunGame()**

## Test ROMs
If you want to test your emulator, here's a link where you can download multiple of them **https://github.com/Mentor-Paris/CHIP-8-Test-Roms**

## Help Code
Not all code has been done by me. I have taken code / inspiration from another project that I have found however I didn't just copy and 
pasted. I made sure to write every line and understand what I'm typing (at least as much I can). To show that, I have written a file
called 'help.txt' where I've explained a few functions that I have took / didn't really understand at first and I thought that it needed
a little description.

## Problems
### Keypad
The keypad can work however you'll need to spam. If you spam it enough, the key will be taken into account and the action should
work (except the 3rd option in the 6-keypad.ch6 test). I could be wrong but I'm feeling that the emulator is somewhat slow. Maybe it
has something to do with the frequency that I have put or it's just bad coded. That could potentially have to do something with the keypad
being irresponsive.

### Register Overflowing
It seems that my emulator is failling pretty bad when it comes to register overflowing. It has failed the 3-corax+.ch test and there's a lot
of problems with the 4-flag.ch test as well.


