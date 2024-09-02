export interface MoodModel { 
    id?:string , 
    name:string , 
    description:string ,
    categoryId:string
    image?:string ,
      }
     // this is to fetch or search
    export interface MoodSearch { 
    
    }
    export interface MoodCategory {
      id?:string ,
      name:string ,
    }
    export const categories:MoodCategory[]=[
      {
        "id": "Positive",
        "name": "Feeling Happy"
      }
      , {
        "id": "27653c90-bee8-4d99-bce5-1854ae6fb09e",
        "name": "Feeling Sad"
      }, 
      {
        "id": "Neutral",
        "name": "Feeling Normal"
      },
      {
        "id": "Complex",
        "name": "Cant Tell ?"
      },{
        "id": "Creative and Productive",
        "name": "Inspired By Someone"
      }
    ]
    export const moods: MoodModel[] = [
        { name: "Happy", description: "Feeling pleasure or contentment.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😊",  },
        { name: "Excited", description: "Feeling enthusiastic and eager.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😃",  },
        { name: "Elated", description: "Extremely happy and excited.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "🤩",  },
        { name: "Content", description: "Satisfied with the current situation.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😌",  },
        { name: "Joyful", description: "Full of joy.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😁",  },
        { name: "Optimistic", description: "Hopeful and confident about the future.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😊",  },
        { name: "Proud", description: "Feeling satisfaction over something regarded as highly honorable or creditable.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😌",  },
        { name: "Relaxed", description: "Free from tension and anxiety.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "😌",  },
        { name: "Grateful", description: "Feeling thankful.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "🙏",  },
        { name: "Loving", description: "Feeling or showing love and affection.", categoryId: "8685c225-dc28-4384-8ebd-0caf3a700b2a", image: "❤️",  },
      
        { name: "Sad", description: "Feeling sorrowful.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😢",  },
        { name: "Angry", description: "Feeling strong annoyance or displeasure.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😠",  },
        { name: "Frustrated", description: "Feeling annoyed or less effective because of a lack of progress.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😤",  },
        { name: "Anxious", description: "Feeling worried or uneasy.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😟",  },
        { name: "Depressed", description: "Feeling severe despondency and dejection.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😞",  },
        { name: "Lonely", description: "Feeling sad because one has no friends or company.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😔",  },
        { name: "Guilty", description: "Feeling responsible for a perceived offense or wrong.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😓",  },
        { name: "Jealous", description: "Feeling or showing envy of someone.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😒",  },
        { name: "Annoyed", description: "Feeling slight anger or irritation.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😒",  },
        { name: "Stressed", description: "Feeling tension and pressure.", categoryId: "27653c90-bee8-4d99-bce5-1854ae6fb09e", image: "😫",  },
      
        { name: "Calm", description: "Peaceful and free from stress.", categoryId: "b5ad3bbc-2757-4331-977a-f5a9ab8591ac", image: "😌",  },
        { name: "Indifferent", description: "Having no particular interest or sympathy; unconcerned.", categoryId: "b5ad3bbc-2757-4331-977a-f5a9ab8591ac", image: "😐",  },
        { name: "Bored", description: "Feeling weary and impatient because one is unoccupied or lacks interest in one's current activity.", categoryId: "b5ad3bbc-2757-4331-977a-f5a9ab8591ac", image: "😒",  },
        { name: "Pensive", description: "Engaged in serious thought.", categoryId: "b5ad3bbc-2757-4331-977a-f5a9ab8591ac", image: "🤔",  },
        { name: "Neutral", description: "Not showing any strong emotion.", categoryId: "b5ad3bbc-2757-4331-977a-f5a9ab8591ac", image: "😐",  },
      
        { name: "Nostalgic", description: "Feeling sentimental longing for the past.", categoryId: "fa2fcb34-ac46-4b43-84f8-a4549348d421", image: "😊",  },
        { name: "Ambivalent", description: "Having mixed feelings or contradictory ideas about something or someone.", categoryId: "fa2fcb34-ac46-4b43-84f8-a4549348d421", image: "😕",  },
        { name: "Conflicted", description: "Feeling uncertain or in a state of conflict.", categoryId: "fa2fcb34-ac46-4b43-84f8-a4549348d421", image: "😖",  },
        { name: "Apprehensive", description: "Anxious or fearful that something bad or unpleasant will happen.", categoryId: "fa2fcb34-ac46-4b43-84f8-a4549348d421", image: "😟",  },
        { name: "Hopeful", description: "Feeling or inspiring optimism about a future event.", categoryId: "fa2fcb34-ac46-4b43-84f8-a4549348d421", image: "😊",  },
      
        { name: "Sociable", description: "Enjoying and seeking the company of others.", categoryId: "e2d123a7-dd81-4384-92fb-64f104ac6389", image: "😃",  },
        { name: "Empathetic", description: "Understanding and sharing the feelings of others.", categoryId: "e2d123a7-dd81-4384-92fb-64f104ac6389", image: "🤗",  },
        { name: "Sympathetic", description: "Feeling compassion for others.", categoryId: "e2d123a7-dd81-4384-92fb-64f104ac6389", image: "😢",  },
      
        { name: "Anticipatory", description: "Looking forward to something with excitement.", categoryId: "7956b74b-f2e1-437d-ae3c-4ab4e8a4dc7e", image: "😃",  },
        { name: "Reflective", description: "Deeply or seriously thoughtful.", categoryId: "7956b74b-f2e1-437d-ae3c-4ab4e8a4dc7e", image: "🤔",  },
        { name: "Reminiscent", description: "Tending to recall or suggest something in the past.", categoryId: "7956b74b-f2e1-437d-ae3c-4ab4e8a4dc7e", image: "😊",  },
      
        { name: "Inspired", description: "Filled with the urge or ability to do or feel something.", categoryId: "b11ef19b-0788-4490-8181-bde18a61c72b", image: "🤩",  },
        { name: "Motivated", description: "Very eager to do or achieve something.", categoryId: "b11ef19b-0788-4490-8181-bde18a61c72b", image: "😃",  },
        { name: "Focused", description: "Having a clear and definite aim or goal.", categoryId: "b11ef19b-0788-4490-8181-bde18a61c72b", image: "😌",  },
      
        { name: "Melancholic", description: "Feeling or expressing pensive sadness.", categoryId: "5de2ee6a-1230-4cb0-a7d0-aa0758a650be", image: "😞",  },
        { name: "Euphoric", description: "Intense excitement and happiness.", categoryId: "5de2ee6a-1230-4cb0-a7d0-aa0758a650be", image: "🤩",  },
        { name: "Determined", description: "Showing firmness of purpose.", categoryId: "5de2ee6a-1230-4cb0-a7d0-aa0758a650be", image: "😠",  },
        { name: "Curious", description: "Eager to know or learn something.", categoryId: "5de2ee6a-1230-4cb0-a7d0-aa0758a650be", image: "🤔",  }
      ];