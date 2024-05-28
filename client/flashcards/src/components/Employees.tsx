import {useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {Card, Client, Department, Employee as EmployeeDTO, Rank} from "../Client"
import {EmployeeCard} from "./EmployeeCard"

const initialCards: EmployeeDTO[] = []
const initialDepartments: Department[] = []
const initialRanks: Rank[] = []
export default function Employees(props: {client: Client}) {
  let client = props.client as Client
  let navigate = useNavigate()
  const [employees, setEmployees] = useState(initialCards)
  const [departments, setDepartments] = useState(initialDepartments)
  const [ranks, setRanks] = useState(initialRanks)

  function deleteEmployee(name: string)  {
    client.deleteEmployee(name).then(() => {
      client.listEmployees().then((list) => {
        setEmployees(list)
      })
    })
  }

  function promoteEmployee(employee_name: string, rank_name: string)  {
    client.promoteEmployee(employee_name, rank_name).then(() => {
      client.listEmployees().then((list) => {
        setEmployees(list)
      })
    })
  }

  function updateFrequency(frequency: number, employee_name: string)  {
    client.updateFrequencyThisMonth(frequency, employee_name).then(() => {
      client.listEmployees().then((list) => {
        setEmployees(list)
      })
    })
  }

  function updateSalary(salary: number, employee_name: string)  {
    client.updateSalary(salary, employee_name).then(() => {
      client.listEmployees().then((list) => {
        setEmployees(list)
      })
    })
  }

  function updateVr(vr: number, employee_name: string)  {
    client.updateVr(vr, employee_name).then(() => {
      client.listEmployees().then((list) => {
        setEmployees(list)
      })
    })
  }

  useEffect(() => {
    if (client.token.Token == "") {
      navigate("/")
      return
    }
    client.listEmployees().then((list) => {
      setEmployees(list)
    })
    client.listDepartments().then((list) => {
      setDepartments(list)
    })

    client.listRanks().then((list) => {
      setRanks(list)
    })
  }, [])

  function listCards() {
    if (employees != null) {
      return employees.map((employee, id) => {
        return (
          <div key={id} style={{marginTop: '70px', maxHeight: 'calc(100vh - 60px)', overflowY: 'auto'}} className="d-flex justify-content-center" >
            <EmployeeCard onUpdateSalary={updateSalary} onUpdateVr={updateVr} key={id} onUpdateFrequency={updateFrequency} client={client} onPromote={promoteEmployee} ranksList={ranks} departmentList={departments} onDelete={deleteEmployee} employee={employee} id={id}></EmployeeCard>
          </div>
        )
      })
    }
  }

  return (
    <>
      {listCards()}
    </>
  )
}

