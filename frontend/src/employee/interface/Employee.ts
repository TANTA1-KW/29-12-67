export interface EmployeesInterface {

    ID?: number;
  
    first_name?: string;
  
    last_name?: string;
  
    email?: string;
  
    phone?: string;
  
    age?: number;

    Address?: string;

    BirthDay?: string;
    
    Password?: string;
    
    salary?: number;

    picture?: string;
    
    gender_id?: number;

    role_id?: number;

    stat_id?: number;

    ship_id?: number;
    
    [x: string]: any;
  }