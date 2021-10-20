# Avalon

## Room State Diagram

```mermaid
stateDiagram-v2
    [*] --> Wait
    Wait --> Team
    Team --> Vote
    Vote --> Quest
    Quest --> Team
    Quest --> [*]
```

## Journey

```mermaid
journey
    title Avalon
    section Join a Room
      Load Web Page: 1: Me
      Identify: 1: Me
      Create a Room: 1: Me
      Join a Room: 1: Me
```

```mermaid
journey
    title Avalon
    section Play a Game
      Is room playing: 1: Me
      Get a character: 1: Me
      Leader Visibility: 1: Me
      See Quests: 1: Me
```
