import React from 'react';
import './ProjectCard.scss'

const ProjectCard = ({ title, time, duration, status }) => {
    return (
        <div className={`project-card ${status}`}>
            <div className="event-info">
                <h3>{title}</h3>
                <p>{time}</p>
            </div>
            <div className="event-meta">
                <span className="duration">{duration}</span>
                <span className={`status-icon ${status}`}></span>
            </div>
        </div>
    );
};

export default ProjectCard;