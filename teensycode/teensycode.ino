const int BUFFER_SIZE = 50;
byte buf[BUFFER_SIZE];
int counter = 0;

void setup() {
  Serial.begin(115200);
}

void loop() {
    // read the incoming bytes:
    int rlen = Serial.readBytes(buf, BUFFER_SIZE);

    // prints the received data
    if (rlen > 0) {
      Serial.print("I received: ");
      for(int i = 0; i < rlen; i++){
        Serial.print(char(buf[i]));
      }
      Serial.println("!");
    }

      Serial.println(counter++);

      delay(500);
}
