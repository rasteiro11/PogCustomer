import {FormEvent, useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {Card, Client, Department, Employee, JwtToken, Rank} from "../Client"

import "../App.css"
import "../Add.css"

const initialDepartments: Department[] = []
const initialEmployees: Employee[] = []
const initialRanks: Rank[] = []
export default function Summary(props: {client: Client}) {
  let client: Client = props.client
  const [departments, setDepartments] = useState(initialDepartments)
  const [ranks, setRanks] = useState(initialRanks)
  const navigate = useNavigate()
  const [employees, setEmployees] = useState(initialEmployees)
  const [salary, setSalary] = useState(0)
  const [frequency, setFrequency] = useState(0)
  const [departmentId, setDepartmentId] = useState(1)
  const [vr, setVr] = useState(0)
  const [formData, setFormData] = useState({
    department_name: 'MARKETING',
    salary: '',
    name: '',
    rank_name: ''
  });

  const getDepartmentName = (list: Department[], id: number): string => {
   const department = departments.find(department => department.ID === id);
     return department ? department.Name : 'Unknown Department';
 };

  useEffect(() => {
    if (client.token.Token == "") {
      navigate("/")
      return
    }
    
      const fetchData = async () => {
        const employeesList = await client.listEmployees()
        const departmentsList = await client.listDepartments()
        const ranksList = await client.listRanks()
        setDepartments(departmentsList)
        console.log("EMPLOYEES LIST: ", employeesList)
        console.log("DEPARTMENTS LIST: ", departmentsList)
        console.log("RANKS LIST: ", ranksList)

        var s = 0
        var v = 0
        var f = 0
        for(var i = 0; i < employeesList.length; i++) {
          if(employeesList[i].department_id == departmentId)  {
            s += employeesList[i].salary
            v += (await client.getVr(employeesList[i].vr_id)).Value
            f += (await client.getFrequencyThisMonth(employeesList[i].name)).TotalDias
          }
        }

        setSalary(s)
        setVr(v)
        setFrequency(f)
      };

      fetchData();
  }, [departmentId])

  const handleChange = (e: any) => {
    const {name, value} = e.target;
    console.log("VALUE: ", value)
    setDepartmentId(value)
    var s = 0
    for(var i = 0; i < employees.length; i++) {
      console.log("DEPARTMENT NAME: ", getDepartmentName(departments, employees[i].department_id))
      console.log("FORMDATA DEPARTMENT NAME: ", formData.department_name)

      console.log("LIST ITEM DEP ID: ", employees[i].department_id)
      if(employees[i].department_id == departmentId)  {
        s += employees[i].salary
      }
    }
    setSalary(s)
    // setFormData({...formData, [name]: value});
  };

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    let loginResponse = await client.addEmployee(formData.department_name, formData.salary, formData.name, formData.rank_name)
    navigate("/employees")
  }

  return <div style={{marginTop: '70px'}} className="d-flex justify-content-center form-container">
    <form onSubmit={handleSubmit}>
      <div>
        <label>Department Name:</label>
        <select name="department_name" value={departmentId} onChange={handleChange}>
          {departments.map(department => (
            <option key={department.ID} value={department.ID}>{department.Name}</option>
          ))}
        </select>
      </div>
      <div>
        <label>Salary: </label>${salary}
      </div>
      <div>
        <label>VR: </label>${vr}
      </div>
      <div>
        <label>Frequency: </label>{frequency}
      </div>
    </form>
  </div>
}
