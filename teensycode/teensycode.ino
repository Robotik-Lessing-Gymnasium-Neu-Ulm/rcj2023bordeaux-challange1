int rx_byte;

void setup() {
  // initialize serial:
  Serial.begin(115200);
}

void loop() {
  // if there's any serial available, read it:
  if (Serial.available() > 0) {
    rx_byte = Serial.read();
        
    Serial.print("=====> ");
    Serial.println(rx_byte);
  }
  Serial.println(rx_byte);
  sendBLTData(rx_byte);
}

void sendBLTData(int& number) {
  if (Serial.availableForWrite()) {
    Serial.write(number);
  }
  
}
