#include "compass.ino"
#include "motors.ino"
#include "bluethooth.ino"

// Gyro libaries
#include <Adafruit_BNO055.h>
#include <Wire.h>
#include <Adafruit_Sensor.h>
#include <EEPROM.h>
#define I2C

// Erstellen eines Obketes der Klasse Adafruit_BNO055 mit Namen gyro
Adafruit_BNO055 gyro = Adafruit_BNO055(55, 0x28);

#define gyroButton 31

#define M1_FW 28 // richtig
#define M1_RW 27
#define M1_PWM 26

#define M2_FW 25 // richtig
#define M2_RW 24
#define M2_PWM 12

#define M3_FW 11 // richtig!S
#define M3_RW 10
#define M3_PWM 9

#define M4_FW 8 // richtig!
#define M4_RW 7
#define M4_PWM 6

#define VR 4 // m1
#define VL 1 // m2
#define HR 3 // m3
#define HL 2 // m4

// Gyro variables
long alteZeit;
int alterWinkel;
double minus;
double rotation;

// Motors
double m1;
double m2;
double m3;
double m4;
double phi;

int Drivedir;

void setup()
{
    Serial.begin(115200);
    Serial1.begin(115200);

    // Linker Motor Pin Festlegung
    pinMode(M1_FW, OUTPUT);
    pinMode(M1_RW, OUTPUT);
    // Hinterer Motor Pin Festlegung
    pinMode(M2_FW, OUTPUT);
    pinMode(M2_RW, OUTPUT);
    // Rechter Motor Pin Festlegung
    pinMode(M3_FW, OUTPUT);
    pinMode(M3_RW, OUTPUT);
    // links hinten
    pinMode(M4_FW, OUTPUT);
    pinMode(M4_RW, OUTPUT);

    pinMode(gyroButton, INPUT_PULLUP);

    gyro.begin();
}

void loop()
{
    compass();
    Drivedir == bluetooth();
    if (!(Drivedir >= 200))
    {
        int redone = Drivedir * 2;
        motor(redone, 50);
    }
}