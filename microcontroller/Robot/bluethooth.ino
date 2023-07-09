int bluetooth() {
    if (Serial5.available()) {
        int angle;
        angle = Serial5.read();
        return angle;
    }    
}