let avalon;

function setup() {
    createCanvas(800, 600, WEBGL);
    ortho();
    angleMode(RADIANS);
    rectMode(CORNER);
    imageMode(CORNER);
    avalon = new Avalon();
}

function draw() {
    avalon.draw();
}

function mouseClicked() {
    console.log(`Mouse clicked! [${mouseX}, ${mouseY}]`);
}
