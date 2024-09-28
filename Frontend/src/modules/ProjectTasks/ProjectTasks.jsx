import React, {useEffect, useState} from 'react';
import Sidebar from "../Sidebar/Sidebar";

import './ProjectTasks.scss'
import Task from "../Task/Task";
import Header from "../Header/Header";
import {useParams} from "react-router-dom";

const ProjectTasks = ({}) => {

    const {id} = useParams()
    const [tasks, setTasks] = useState([])
    const [activeTask, setActiveTask] = useState({})


    useEffect(() => {
        fetch('https://a69e816f684673.lhr.life/v1/projects/H%26h/tasks')
            .then((response) => response.json())
            .then((json) => {
                setTasks(json.data)
            })
            .catch((err) => {
                console.warn(err)
            })
    }, [])


    return (
        <div className="tasks-container">
            <Sidebar/>
            <main className="content">
                <Header title='Project Tasks' buttonText='Add tasks'/>
                <section className="task-details">
                    <div className="task-list">
                        <h3>Current Tasks</h3>
                        <ul>
                            {tasks.filter((task) => task.project_id === parseInt(id))
                                .map((task, index) => (
                                    <li key={index}>{task.Title}</li>
                                ))
                            }
                        </ul>
                    </div>
                    <Task/>
                </section>
            </main>
        </div>
    );
};

export default ProjectTasks;