let avalon

let FONT_INCONSOLATA

let IMAGE_TITLE
let IMAGE_CHARACTER_CARD_BACK
let IMAGE_ASSASSIN
let IMAGE_MERLIN
let IMAGE_MINION0
let IMAGE_MINION1
let IMAGE_MINION2
let IMAGE_MORGANA
let IMAGE_OBERON
let IMAGE_PERCIVAL
let IMAGE_SERVANT0
let IMAGE_SERVANT1
let IMAGE_SERVANT2
let IMAGE_SERVANT3
let IMAGE_SERVANT4

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
    IMAGE_CHARACTER_CARD_BACK = loadImage("images/character_back.png")
    IMAGE_ASSASSIN = loadImage("images/character_assassin.png")
    IMAGE_MERLIN = loadImage("images/character_merlin.png")
    IMAGE_MINION0 = loadImage("images/character_minion0.png")
    IMAGE_MINION1 = loadImage("images/character_minion1.png")
    IMAGE_MINION2 = loadImage("images/character_minion2.png")
    IMAGE_MORGANA = loadImage("images/character_morgana.png")
    IMAGE_OBERON = loadImage("images/character_oberon.png")
    IMAGE_PERCIVAL = loadImage("images/character_percival.png")
    IMAGE_SERVANT0 = loadImage("images/character_servant0.png")
    IMAGE_SERVANT1 = loadImage("images/character_servant1.png")
    IMAGE_SERVANT2 = loadImage("images/character_servant2.png")
    IMAGE_SERVANT3 = loadImage("images/character_servant3.png")
    IMAGE_SERVANT4 = loadImage("images/character_servant4.png")

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

function mouseClicked() {
    console.log("Mouse Clicked")
    avalon.merlin_card.checkBeingClicked()
    avalon.morgana_card.checkBeingClicked()
}

function mousePressed() {
    console.log("Mouse Pressed")
    console.log(`${mouseX} ${mouseY} ${avalon.menu.createRoomButton.bbox.x} ${avalon.menu.createRoomButton.bbox.y}`)
}

function mouseDragged() {
    console.log("Mouse Dragged")
  }

  function mouseReleased() {
    console.log("Mouse Released")
  }
