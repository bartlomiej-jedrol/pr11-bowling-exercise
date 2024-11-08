# Bowling Score calculator

Write a program that scores bowling.

## Scoring Rules
 
The game consists of 10 frames as shown below.  In each frame the player has two opportunities to knock down 10 pins.  The score for the frame is the total number of pins knocked down, plus any bonuses for strikes and spares.

A spare is when the player knocks down all 10 pins in two tries.  The bonus for that frame is the number of pins knocked down by the next roll.  So in frame 3 below, the score is 10 (the total number knocked down) plus a bonus of 5 (the number of pins knocked down on the next roll) for a total of 15.  

A strike is when the player knocks down all 10 pins on his first try.  The bonus for that frame is the value of the next two balls rolled.  So in frame 5 below, the score is 10 plus bonuses of 0 and 1 (the number of pins knocked down on the next two rolls) for a total of 11.

In the tenth frame a player who rolls a spare or strike is allowed to roll the extra balls to complete the frame.  However no more than three balls can be rolled in tenth frame, so any strikes in the bonus rolls do not also earn bonus rolls.

![Bowling Scorecard](assets/exercise.png)

## Simplified Scoring Rules

1. **Game Structure**
   - Game consists of 10 frames
   - Each frame allows up to 2 rolls to knock down 10 pins
   - Frame score = pins knocked down + bonus points (if any)

2. **Spare** (/)
   - All 10 pins knocked down in 2 rolls
   - Bonus: Points from next roll
   - Example: Frame score = 10 + next roll

3. **Strike** (X)
   - All 10 pins knocked down in first roll
   - Bonus: Points from next two rolls
   - Example: Frame score = 10 + next two rolls

4. **Tenth Frame Special Rules**
   - Spare: One bonus roll allowed
   - Strike: Two bonus rolls allowed
   - Maximum three rolls total
   - Bonus rolls only count for points (no additional bonus rolls)

## Solutions

### Solution 1 - Random Bowling Game Score Generator API

Random Bowling Game Score Generator API
- The program handles a single bowling game
- The program generates random roll scores between 1 and 10
- The frame score does not include bonuses for spare or strike, as these are included only in the final game score
- The program is hosted on AWS Lambda and exposed as a public API through AWS API Gateway
- The API returns the bowling game result in raw JSON format (sample JSON included below)
- The program is available at the API endpoint: https://hljvgios4m.execute-api.eu-central-1.amazonaws.com/dev
- The source code for the solution is available in the following directories:
   - bowling/bowling.go
   - cmd/lambda/main.go

#### Sample JSON response
```JSON
{
  "frames": [
    {
      "rolls": [6, 0],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 6
    },
    {
      "rolls": [6, 1],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 7
    },
    {
      "rolls": [9, 0],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 9
    },
    {
      "rolls": [10],
      "is_strike": true,
      "is_spare": false,
      "frame_score": 10
    },
    {
      "rolls": [8, 0],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 8
    },
    {
      "rolls": [0, 5],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 5
    },
    {
      "rolls": [6, 4],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 10
    },
    {
      "rolls": [8, 2],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 10
    },
    {
      "rolls": [5, 5],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 10
    },
    {
      "rolls": [1, 9, 8],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 18
    }
  ],
  "game_score": 115
}
```

#### Sample JSON reponse visualisation
![Random Bowling Game Score Generator sample](assets/random_bowling_game_score_generator_sample.png)
   
### Solution 2 - Bowling Game Score CLI

Bowling Game Score CLI
- The program handles a single bowling game
- The program gets input on rolls from the user
- The frame score includes bonuses for spares and strikes
- The program is built for the following operating systems:
   - Windows
   - Linux
   - Mac
- The programs for all operating systems are available at the following link: https://drive.google.com/drive/folders/1i-UowsTGLMsmk2rofHkrhMHEUTO_UYCu?usp=sharing
- The source code for the solution is available in the following directory:
   - cmd/cli/main.go

#### Sample Bowling Game Score CLI game
![Bowling Game Score CLI](assets/bowling_game_score_cli_sample.png)