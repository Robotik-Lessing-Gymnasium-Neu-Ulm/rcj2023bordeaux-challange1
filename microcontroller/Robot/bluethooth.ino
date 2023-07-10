int bluetooth() {
    if (Serial1.available()) {
        int angle;
        angle = Serial1.read();
        return angle;
    }    
}