export interface IAhp {
  goal_name:string;
  alternatives:Array<IAlternative>;
  criteria:Array<ICriteria>;
}

interface IAlternative {
  name:string;
  final_score?:number;
}

interface ICriteria {
  name:string;
  weight:number;
  subcriteria?:Array<ICriteria>;
  weights?:Array<number>;
}


var ahp = {
  goal: 'car',
  alternatives: [
    {name: 'car1', final_score: null},
    {name: 'car2', final_score: null},
    {name: 'car3', final_score: null}
  ],
  criteria: [
    {
      name: 'engine',
      weight: 0.45,
      subcriteria: [
        {
          name: 'horsepower',
          weight: 0.5,
          weights: [0.3, 0.3, 0.4]
        },
        {
          name: 'durability',
          weight: 0.5,
          weights: [0.3, 0.3, 0.4]
        }
      ]
    },
    {
      name: 'design',
      weight: 0.65,
      weights: [0.4, 0.4, 0.2]
    }
  ]
}
