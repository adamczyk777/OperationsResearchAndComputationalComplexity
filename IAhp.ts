export interface IAhp {
  goal_name:string;
  alternatives:Array<IAlternative>;
  criteria:Array<ICriteria>;
}

interface IAlternative {
  name:string;
  score?:number;
}

interface ICriteria {
  name:string;
  score_sum:number;
  subcriteria:Array<ICriteria>;
}
