import React, { useState, useEffect } from 'react';
import './Projects.scss';
import ProjectCard from "../ProjectCard/ProjectCard";
import Sidebar from "../Sidebar/Sidebar";
import Header from '../Header/Header';
import { useNavigate } from "react-router-dom";
import Loader from "../../assets/Loader/Loader";

const Projects = () => {
    const navigate = useNavigate();
    const [projects, setProjects] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null); 

    const onClickProjectDetails = (id) => {
        navigate(`/project/${id}`);
    };

    useEffect(() => {
        fetch('https://a69e816f684673.lhr.life/v1/projects')
            .then((response) => {
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }
                return response.json();
            })
            .then((json) => {
                setProjects(json.data || []);
                setLoading(false);
            })
            .catch((err) => {
                console.error('Error fetching projects:', err);
                setError('Failed to load projects');
                setLoading(false);
            });
    }, []);

    return (
        <div className="dashboard">
            <Sidebar />
            <div className="content">
                <Header title="Projects" buttonText="Add Project" isAdd={true} isSearched={true} />
                <section className="nearest-events">
                    {loading ? (
                        <div className='flex justify-center items-center mt-20 flex-col'>
                            <Loader />
                            <h1 className="text-2xl font-bold text-gray-800">Loading...</h1>
                        </div>
                    ) : error ? (
                        <div className='flex justify-center items-center mt-20 flex-col'>
                            <h1 className="text-2xl font-bold text-red-600">{error}</h1>
                        </div>
                    ) : projects.length ? (
                        <div className="events-list">
                            {projects.map((obj, index) => (
                                <ProjectCard
                                    key={obj.id || index}
                                    {...obj}
                                    onClickProjectDetails={onClickProjectDetails}
                                />
                            ))}
                        </div>
                    ) : (
                        <div className='flex justify-center items-center mt-20 flex-col'>
                            <h1 className="text-2xl font-bold text-gray-800">No projects found :(</h1>
                        </div>
                    )}
                </section>
            </div>
        </div>
    );
};

export default Projects;
