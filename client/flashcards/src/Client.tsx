export interface Employee {
  id: number
  name: string
  department_id: number
  vr_id: number
  salary: number
  rank_id: number
}

export interface Vr {
  ID: number
  Value: number
}

export interface Frequency {
  DataReferencia: Date
  TotalDias: number
}

export interface Rank {
  ID: number
  Name: string
}

export interface Department {
  ID: number
  Name: string
}

export interface Cards {
  Data: any[]
}

export class Card {
  ID = 0
  CreatedAt = new Date
  UpdatedAt = new Date
  UserID = 0
  WhichBox: number
  Question: string
  Answer: string

  constructor(box: number, question: string, answer: string) {
    this.WhichBox = box
    this.Question = question
    this.Answer = answer
  }
}

export class JwtToken {
  Token: string
  Expiration: Date

  constructor() {
    this.Token = ""
    this.Expiration = new Date()
  }
}

function getFirstDayOfCurrentMonth() {
    const now = new Date();
    const firstDay = new Date(now.getFullYear(), now.getMonth(), 1);

    const year = firstDay.getFullYear();
    const month = String(firstDay.getMonth() + 1).padStart(2, '0');
    const day = String(firstDay.getDate()).padStart(2, '0');
    const hours = '00';
    const minutes = '00';
    const seconds = '00';

    const timezoneOffset = -firstDay.getTimezoneOffset();
    const absOffsetHours = String(Math.floor(Math.abs(timezoneOffset) / 60)).padStart(2, '0');
    const absOffsetMinutes = String(Math.abs(timezoneOffset) % 60).padStart(2, '0');
    const sign = timezoneOffset >= 0 ? '+' : '-';

    const formattedDate = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}${sign}${absOffsetHours}:${absOffsetMinutes}`;
    return formattedDate;
}

var baseCustomerUrl = "http://localhost:8080/customer"
var basePaymentSheetUrl = "http://localhost:8081/payment/sheet"
export class Client {
  token: JwtToken
  constructor() {
    this.token = new JwtToken()
  }

  async addEmployee(department_name: string,
    salary: string,
    name: string,
    rank_name: string) {

      return call(basePaymentSheetUrl+"/employee/add", "POST", {
        department_name: department_name,
        salary: parseFloat(salary),
        name: name,
        rank_name: rank_name
      },{
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }

  async updateFrequencyThisMonth(frequency: number, employeeName: string): Promise<Frequency> {
    return call(basePaymentSheetUrl+"/frequency/update", "POST", {
      Date: getFirstDayOfCurrentMonth(),
      EmployeeName: employeeName,
      Frequency: frequency
    }, {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async getFrequencyThisMonth(employeeName: string): Promise<Frequency> {
    return call(basePaymentSheetUrl+"/frequency/get", "POST", {
      Date: getFirstDayOfCurrentMonth(),
      EmployeeName: employeeName
    }, {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async updateSalary(vr: number, employeeName: string): Promise<Vr> {
    return call(basePaymentSheetUrl+"/employee/update", "POST", {
      EmployeeName:employeeName,
      Value: vr
    }, {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async updateVr(vr: number, employeeName: string): Promise<Vr> {
    return call(basePaymentSheetUrl+"/vr/update", "POST", {
      EmployeeName:employeeName,
      Value: vr
    }, {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async getVr(vrId: number): Promise<Vr> {
    return call(basePaymentSheetUrl+"/vr/find", "POST", {
      vr_id: vrId
    }, {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async listRanks(): Promise<Rank[]> {
    return call(basePaymentSheetUrl+"/rank/list", "POST", {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }


  async listDepartments(): Promise<Department[]> {
    return call(basePaymentSheetUrl+"/department/list", "POST",
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }

  async findOneCard(id: number): Promise<Card> {
    let idStr: string = id.toString()
    return call("/user/flashcard/"+idStr, "GET",
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async listCards(): Promise<Cards> {
    return {Data: []}
  }

  async deleteCard(id: number): Promise<Card> {
    let idStr: string = id.toString()
    return call("/user/flashcard/"+idStr, "DELETE", 
      {},
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }

  async promoteEmployee(employee_name: string, rank_name: string) {
    return call(basePaymentSheetUrl+"/employee/promote", "POST",
      {
        employee_name: employee_name,
        rank_name: rank_name
      },
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }

  async deleteEmployee(name: string)  {
    return call(basePaymentSheetUrl+"/employee/delete", "POST",
      {
        employee_name: name
      },
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async listEmployees(): Promise<Employee[]> {
    return call(basePaymentSheetUrl+"/employee/list", "POST",
      {
      },
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })

  }

  async createCard(card: Card): Promise<Card> {
    return call("/user/flashcard", "POST", card,
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }

  async register(email: string, password: string): Promise<JwtToken> {
    return call(baseCustomerUrl+"/auth/register", "POST",
      {
        "email": email,
        "password": password
      },
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }

  async login(email: string, password: string): Promise<JwtToken> {
    return call(baseCustomerUrl+"/auth/signin", "POST",
      {
        "email": email,
        "password": password
      },
      {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        'Authorization': "Bearer " + this.token.Token
      })
  }
}

async function call<B, R>(url: string, method: string, body?: B, header?: {[key: string]: string}): Promise<R> {
  let reqConfig: RequestInit = {
    method: method,
    headers: header,
    body: JSON.stringify(body)
  }

  if (body !== undefined) {
    reqConfig.body = JSON.stringify(body)
  }

  let req = await fetch(url, reqConfig)
  return req.json() as Promise<R>
}



