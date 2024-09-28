import React, {useEffect} from 'react';
import './Projects.scss'
import { FiPlus } from 'react-icons/fi';

import ProjectCard from "../ProjectCard/ProjectCard";
import Sidebar from "../Sidebar/Sidebar";
import Header from '../Header/Header';
import {useNavigate} from "react-router-dom";

const Projects = () => {

    const navigate = useNavigate()

    const projectDetails  = (id) => {
        navigate(`/project/${id}`)
    }





    return (
        <div className="dashboard">
            <Sidebar/>
            <div className="content">
                <Header title="Projects" buttonText="Add Project"/>
                <section className="nearest-events">
                    <div className="events-list">
                        <ProjectCard
                            title="Presentation of the new department"
                            time="Today | 6:00 PM"
                            duration="4h"
                            status="up"
                        />
                        <ProjectCard
                            title="Anna's Birthday"
                            time="Today | 5:00 PM"
                            duration="2h"
                            status="down"
                        />
                        <ProjectCard
                            title="Meeting with Development Team"
                            time="Tomorrow | 5:00 PM"
                            duration="4h"
                            status="up"
                        />
                        <ProjectCard
                            title="Ray's Birthday"
                            time="Tomorrow | 2:00 PM"
                            duration="1h 30m"
                            status="down"
                        />
                        <ProjectCard
                            title="Meeting with CEO"
                            time="Sep 14 | 5:00 PM"
                            duration="1h"
                            status="up"
                        />
                        <ProjectCard
                            title="Movie night (Tenet)"
                            time="Sep 15 | 5:00 PM"
                            duration="3h"
                            status="down"
                        />
                        <ProjectCard
                            title="Lucas's Birthday"
                            time="Sep 29 | 5:30 PM"
                            duration="2h"
                            status="down"
                        />
                        <ProjectCard
                            title="Meeting with CTO"
                            time="Sep 30 | 12:00 PM"
                            duration="1h"
                            status="up"
                        />
                    </div>
                </section>
            </div>
        </div>
    );
};


export default Projects;