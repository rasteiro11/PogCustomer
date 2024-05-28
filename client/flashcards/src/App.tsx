import {useState} from 'react';
import {Link, Route, Routes} from 'react-router-dom';
import {Client} from './Client';
import Employees from './components/Employees';
import Login from './components/Login';
import Register from './components/Register';
import Add from './components/Add';
import Summary from './components/Summary';


function Header(props: {loggedIn: boolean}) {
  if (props.loggedIn) {
    return (
      <nav style={
        {position: 'fixed', top: '0', width: '100%', zIndex: 1000}
      } className="navbar navbar-expand-lg navbar-light bg-light">
        <a className="navbar-brand" href="#">Payment Sheet</a>
        <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarNav">
          <ul className="navbar-nav">
            <li className="nav-item"><Link className="nav-link" to='/employees'>Your Employees</Link></li>
            <li className="nav-item"><Link className="nav-link" to='/employee/add'>Add Employee</Link></li>
            <li className="nav-item"><Link className="nav-link" to='/employee/summary'>Payment Sheet Summary</Link></li>
          </ul>
        </div>
      </nav>
    )
  }
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <a className="navbar-brand" href="#">Payment Sheet</a>
      <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span className="navbar-toggler-icon"></span>
      </button>
      <div className="collapse navbar-collapse" id="navbarNav">
        <ul className="navbar-nav">
          <li className='nav-item'><Link  className="nav-link" to='/'>Login</Link></li>
        </ul>
      </div>
    </nav>
  )
}


var client = new Client()
export default function App() {
  const [loggedIn, setLoggedIn] = useState(false)
  return (
    <>
        <Header loggedIn={loggedIn} />
        <hr style={{margin: '0px'}}/>
        <Routes>
          <Route path='/' element={<Login client={client} setLoggedInFn={setLoggedIn} />} />
          <Route path='/register' element={<Register client={client} setLoggedInFn={setLoggedIn} />} />
          <Route path='/employees' element={<Employees client={client} />} />
          <Route path='/employee/add' element={<Add client={client} />} />
          <Route path='/employee/summary' element={<Summary client={client} />} />
        </Routes>
    </>
  )
}


