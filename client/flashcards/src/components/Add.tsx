import {FormEvent, useEffect, useState} from "react"
import {useNavigate} from "react-router-dom"
import {Card, Client, Department, Employee, JwtToken, Rank} from "../Client"

import "../App.css"
import "../Add.css"

const initialDepartments: Department[] = []
const initialRanks: Rank[] = []
export default function Add(props: {client: Client}) {
  let client: Client = props.client
  const [departments, setDepartments] = useState(initialDepartments)
  const [ranks, setRanks] = useState(initialRanks)
  const navigate = useNavigate()
  const [formData, setFormData] = useState({
    department_name: '',
    salary: '',
    name: '',
    rank_name: ''
  });

  useEffect(() => {
    if (client.token.Token == "") {
      navigate("/")
      return
    }
    client.listDepartments().then((list) => {
      setDepartments(list)
    })

    client.listRanks().then((list) => {
      setRanks(list)
    })
  }, [])

  const handleChange = (e: any) => {
    const {name, value} = e.target;
    setFormData({...formData, [name]: value});
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
        <select name="department_name" value={formData.department_name} onChange={handleChange}>
          <option value="">Select Department</option>
          {departments.map(department => (
            <option key={department.ID} value={department.Name}>{department.Name}</option>
          ))}
        </select>
      </div>
      <div>
        <label>Salary:</label>
        <input type="number" name="salary" value={formData.salary} onChange={handleChange} />
      </div>
      <div>
        <label>Name:</label>
        <input type="text" name="name" value={formData.name} onChange={handleChange} />
      </div>
      <div>
        <label>Rank Name:</label>
        <select name="rank_name" value={formData.rank_name} onChange={handleChange}>
          <option value="">Select Rank</option>
          {ranks.map(rank => (
            <option key={rank.ID} value={rank.Name}>{rank.Name}</option>
          ))}
        </select>
      </div>
      <button type="submit">Submit</button>
    </form>
  </div>
}
