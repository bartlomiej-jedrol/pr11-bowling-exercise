# Bowling Score calculator

Write a program that scores bowling.

## Scoring Rules
 
The game consists of 10 frames as shown below.  In each frame the player has two opportunities to knock down 10 pins.  The score for the frame is the total number of pins knocked down, plus any bonuses for strikes and spares.

A spare is when the player knocks down all 10 pins in two tries.  The bonus for that frame is the number of pins knocked down by the next roll.  So in frame 3 below, the score is 10 (the total number knocked down) plus a bonus of 5 (the number of pins knocked down on the next roll) for a total of 15.  

A strike is when the player knocks down all 10 pins on his first try.  The bonus for that frame is the value of the next two balls rolled.  So in frame 5 below, the score is 10 plus bonuses of 0 and 1 (the number of pins knocked down on the next two rolls) for a total of 11.

In the tenth frame a player who rolls a spare or strike is allowed to roll the extra balls to complete the frame.  However no more than three balls can be rolled in tenth frame, so any strikes in the bonus rolls do not also earn bonus rolls.

![Bowling Scorecard](image.png)

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

I prepared two solutions for this exercise:

1. Random Bowling Game Score Generator
- The program handles a single bowling game
- The program generates random roll scores between 1 and 10
- The frame scores do not include bonuses for spare or strike as these are calculated in the final game score
- The program is hosted on AWS Lambda and exposed as a public API through AWS API Gateway
- The API returns the bowling game result in raw JSON format (sample JSON attached in the ...)
- The program is available at the API endpoint: https://hljvgios4m.execute-api.eu-central-1.amazonaws.com/dev

Sample response:
```JSON
{
  "game_id": 1,
  "frames": [
    {
      "frame_id": 1,
      "rolls": [
        {
          "roll_id": 1,
          "roll_score": 2
        },
        {
          "roll_id": 2,
          "roll_score": 3
        }
      ],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 5
    },
    {
      "frame_id": 2,
      "rolls": [
        {
          "roll_id": 3,
          "roll_score": 7
        },
        {
          "roll_id": 4,
          "roll_score": 1
        }
      ],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 8
    },
    {
      "frame_id": 3,
      "rolls": [
        {
          "roll_id": 5,
          "roll_score": 6
        },
        {
          "roll_id": 6,
          "roll_score": 3
        }
      ],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 9
    },
    {
      "frame_id": 4,
      "rolls": [
        {
          "roll_id": 7,
          "roll_score": 1
        },
        {
          "roll_id": 8,
          "roll_score": 7
        }
      ],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 8
    },
    {
      "frame_id": 5,
      "rolls": [
        {
          "roll_id": 9,
          "roll_score": 4
        },
        {
          "roll_id": 10,
          "roll_score": 1
        }
      ],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 5
    },
    {
      "frame_id": 6,
      "rolls": [
        {
          "roll_id": 11,
          "roll_score": 10
        }
      ],
      "is_strike": true,
      "is_spare": false,
      "frame_score": 10
    },
    {
      "frame_id": 7,
      "rolls": [
        {
          "roll_id": 12,
          "roll_score": 7
        },
        {
          "roll_id": 13,
          "roll_score": 3
        }
      ],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 10
    },
    {
      "frame_id": 8,
      "rolls": [
        {
          "roll_id": 14,
          "roll_score": 0
        },
        {
          "roll_id": 15,
          "roll_score": 6
        }
      ],
      "is_strike": false,
      "is_spare": false,
      "frame_score": 6
    },
    {
      "frame_id": 9,
      "rolls": [
        {
          "roll_id": 16,
          "roll_score": 7
        },
        {
          "roll_id": 17,
          "roll_score": 3
        }
      ],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 10
    },
    {
      "frame_id": 10,
      "rolls": [
        {
          "roll_id": 18,
          "roll_score": 9
        },
        {
          "roll_id": 19,
          "roll_score": 1
        },
        {
          "roll_id": 20,
          "roll_score": 1
        }
      ],
      "is_strike": false,
      "is_spare": true,
      "frame_score": 11
    }
  ],
  "game_score": 101
} 
```

![Random Bowling Game Score Generator sample](Random_Bowling_Game_Score_Generator_sample.png)
   
