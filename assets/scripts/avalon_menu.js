class AvalonMenu {
    constructor() {
        let menu_background = floor(random(14));
        this.image_menu_background = loadImage(`images/backgrounds/background${menu_background}.jpg`, 800, 600);
        this.image_players = loadImage(`images/text/players.png`);
        this.title = new AvalonText("images/text/title.png");
        this.image_numbers = loadImage(`images/numbers.png`);
        this.one = {
            ox: 0,
            oy: 0,
            width: 203,
            height: 211
        }
    }
    draw() {
        push();
        imageMode(CENTER);
        image(this.image_menu_background, 0, 0);

        this.title.draw();

        push();
        imageMode(CORNER);
        stroke(0);
        strokeWeight(5);
        fill(255,255,255,255);
        translate(-200, -100);
        rect(0,0,150,110);
        image(this.image_numbers, 0, 0,
            this.one.width * 0.5, this.one.height * 0.5,
            this.one.ox, this.one.oy,
            this.one.width, this.one.height);
        pop();

        pop();
    }
}
