import {Route, Routes, BrowserRouter} from "react-router-dom";
import SignIn from './modules/SingIn/SignIn';
import SignUp from './modules/SignUp/SignUp';
import Projects from './modules/Projects/Projects';
import Dashboard from "./modules/Dashboard/Dashboard";
import Profile from "./modules/Profile/Profile";
import ProjectTasks from "./modules/ProjectTasks/ProjectTasks";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/login" element={<SignIn/>}/>
                <Route path="/register" element={<SignUp/>}/>
                <Route path="/" element={<Projects/>}/>
                <Route path='/project/:id' element={<ProjectTasks/>}/>
                <Route path="/dashboard" element={<Dashboard/>}/>
                <Route path="/profile" element={<Profile/>}/>
            </Routes>
        </BrowserRouter>

    );
}

export default App;