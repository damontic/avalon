class Avalon {
    constructor(){
        this.state = "menu";
        this.menu = new AvalonMenu();
    }
    draw() {
        switch(this.state) {
            case "menu":
                this.menu.draw();
                break;
            default:
                break;
        }
        // push();
        // let morderedWidth = 91;
        // let morderedHeight = 152;
        // translate(-morderedWidth/2 -windowWidth/4, -morderedHeight/2 -windowHeight/4);
        // image(this.avalonImages, 100, 100, morderedWidth, morderedHeight, 
        //     24, 182, morderedWidth, morderedHeight);
        // pop();
    }
}

class UnsupportedNumberOfPlayers extends Error {
    constructor(...params) {
        super(...params);
        this.name = "UnsupportedNumberOfPlayers";
    }
}
