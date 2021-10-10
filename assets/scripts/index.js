let avalon

let FONT_INCONSOLATA

let IMAGE_TITLE

let RED
let GREEN
let BLUE
let YELLOW
let CYAN
let MAGENTA
let BLACK
let WHITE

function preload() {
    FONT_INCONSOLATA = loadFont('fonts/InconsolataN-Regular.otf')

    IMAGE_TITLE = loadImage("images/text/title.png")

    RED = color(255, 0, 0)
    GREEN = color(0, 255, 0)
    BLUE = color(0, 0, 255)
    YELLOW = color(255, 255, 0)
    MAGENTA = color(255, 0, 255)
    CYAN = color(0, 255, 255)
    BLACK = color(0)
    WHITE = color(255)
}

function setup() {
    let appDisplayWidth = displayWidth * 80 / 100
    let appDisplayHeight = displayHeight * 80 / 100
    createCanvas(appDisplayWidth, appDisplayHeight, WEBGL)
    ortho()
    angleMode(RADIANS)
    rectMode(CORNER)
    textAlign(LEFT, TOP)
    avalon = new Avalon(
        appDisplayWidth,
        appDisplayHeight,
        getURLParams()
    )
}

function draw() {
    avalon.draw()
}

function mousePressed() {
    console.log(`${mouseX} ${mouseY} ${avalon.menu.createRoomButton.bbox.x} ${avalon.menu.createRoomButton.bbox.y}`)
}
