import {useEffect, useState} from "react";
import "../App.css"
import {Client, Department, Employee, Rank} from "../Client";
import '../EmployeeCard.css';

export function EmployeeCard(props: {client: Client, ranksList: Rank[], departmentList: Department[], employee: Employee, id: number,
                             onUpdateFrequency: (freq: number, employeeName: string) => void,
                             onUpdateVr: (vr: number, employeeName: string) => void,
                             onUpdateSalary: (salary: number, employeeName: string) => void,
                             onDelete: (name: string) => void,
                             onPromote: (employee_name: string, rank_name: string) => void
    }) {
    const getRankName = (id: number): string => {
        const rank = props.ranksList.find(rank => rank.ID === id);
        return rank ? rank.Name : 'Unknown Rank';
    };

    const [vr, setVr] = useState<number>(0);
    const [initialVr, setInitialVr] = useState<number>(0);
    const [initialSalary, setInitialSalary] = useState<number>(0);
    const [salary, setSalary] = useState<number>(0);

    const [selectedRank, setSelectedRank] = useState<string>();

    const [frequency, setFrequency] = useState<number>(0);
    const [initialFrequency, setInitialFrequency] = useState<number>(0);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await props.client.getVr(props.employee.vr_id)
                setVr(response.Value);
                setInitialVr(response.Value)
            } catch (error) {
                console.error('Error fetching data:', error);
            }

            try {
                const response = await props.client.getFrequencyThisMonth(props.employee.name)
                setFrequency(response.TotalDias);
                setInitialFrequency(response.TotalDias)
            } catch (error) {
                console.error('Error fetching data:', error);
            }

            setSelectedRank(getRankName(props.employee.rank_id))
            setInitialSalary(props.employee.salary)
            setSalary(props.employee.salary)
        };

        fetchData();


    }, [])

    const onUpdateFrequencyWrapper = async (frequency: number, name: string) => {
        props.onUpdateFrequency(frequency, props.employee.name)
        setFrequency(frequency)
        setInitialFrequency(frequency)
    }

    const onUpdateVrWrapper = async (vr: number, name: string) => {
        props.onUpdateVr(vr, props.employee.name)
        setVr(vr)
        setInitialVr(vr)
    }

    const onUpdateSalaryWrapper = async (salary: number, name: string) => {
        props.onUpdateSalary(salary, props.employee.name)
        setInitialSalary(salary)
        setSalary(salary)
    }

    const getDepartmentName = (id: number): string => {
        const department = props.departmentList.find(dept => dept.ID === id);
        return department ? department.Name : 'Unknown Department';
    };

    const handleRankChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
        setSelectedRank(event.target.value);
    };

    const handleFrequencyChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        var f = parseInt(event.target.value, 10)
        setFrequency(f);
    };

    const handleVrChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        var f = parseInt(event.target.value, 10)
        setVr(f);
    };

    const handleSalaryChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        var f = parseInt(event.target.value, 10)
        setSalary(f);
    };

    return (
        <div className="employee-card">
            <h2>{props.employee.name}</h2>
            <p><strong>ID:</strong> {props.employee.id}</p>
            <p><strong>Department:</strong> {getDepartmentName(props.employee.department_id)}</p>
            <p><strong>VR:</strong> ${vr}</p>
            <p><strong>Salary:</strong> ${salary}</p>
            <p><strong>Rank:</strong> {getRankName(props.employee.rank_id)}</p>
            <p><strong>Employee Frequency:</strong> {frequency}</p>
            <div className="input-container">
                <label>Salary:</label>
                <input type="number" value={salary} onChange={handleSalaryChange} />
            </div>
            <div className="input-container">
                <label>Employee Frequency:</label>
                <input type="number" value={frequency} onChange={handleFrequencyChange} />
            </div>
            <div className="input-container">
                <label>Change Rank:</label>
                <select value={selectedRank} onChange={handleRankChange}>
                    {props.ranksList.map(rank => (
                        <option key={rank.Name} value={rank.Name}>{rank.Name}</option>
                    ))}
                </select>
            </div>
            <div className="input-container">
                <label>Vr Value:</label>
                <input type="number" value={vr} onChange={handleVrChange} />
            </div>
            <div className="button-container">
                {salary !== initialSalary ?
                    <button onClick={() => onUpdateSalaryWrapper(salary, props.employee.name)}>Confirm Salary Change</button> : <></>}
                {vr !== initialVr ?
                    <button onClick={() => onUpdateVrWrapper(vr, props.employee.name)}>Confirm Vr Change</button> : <></>}
                {frequency !== initialFrequency ?
                    <button onClick={() => onUpdateFrequencyWrapper(frequency, props.employee.name)}>Confirm Frequency Change</button> : <></>}
                {(getRankName(props.employee.rank_id) !== selectedRank && selectedRank !== 'Unknown Rank') ?
                    <button onClick={() => props.onPromote(props.employee.name, selectedRank!)}>Confirm Rank Change</button> : <></>}
                <button className="delete-button" onClick={() => props.onDelete(props.employee.name)}>Delete</button>
            </div>
        </div>
    );
}

