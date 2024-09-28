import React, {useEffect, useState} from 'react';
import Sidebar from "../Sidebar/Sidebar";

import './ProjectTasks.scss'
import Task from "../Task/Task";
import Header from "../Header/Header";
import {json, useNavigate} from "react-router-dom";

const ProjectTasks = () => {




    const [tasks, setTasks] = useState([])


    // useEffect(() => {
    //     fetch('https://69fbdebe0a2299.lhr.life/v1/projects')
    //         .then((response) => response.json())
    //         .then(json => {
    //             setTasks(json)
    //         })
    // }, [])
    //
    // console.log(tasks.data)




    return (
        <div className="tasks-container">
            <Sidebar/>
            <main className="content">
                <Header/>
                <header className="header">
                    <input type="search" placeholder="Search" className="search-bar"/>
                    <button className="add-task-btn">+ Add Task</button>
                </header>
                <section className="task-details">
                    <div className="task-list">
                        <h3>Current Tasks</h3>
                        <ul>
                            <li>Medical App (iOS native)</li>
                            <li>Food Delivery Service</li>
                            <li>Fortune Website</li>
                            <li>Planner App</li>
                            <li>Time Tracker - Personal Account</li>
                            <li>Internal Project</li>
                        </ul>
                    </div>
                    <Task/>
                </section>
            </main>
        </div>
    );
};

export default ProjectTasks;