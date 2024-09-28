import React from 'react';
import './ProjectCard.scss'

const ProjectCard = ({CreatedAt, status, Name, Id, onClickProjectDetails }) => {

    const formattedDate = new Date(CreatedAt).toLocaleDateString("ru-RU", {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
    })


    return (
        <div className={`project-card ${status}`} onClick={() => onClickProjectDetails(Id)}>
            <div className="event-info">
                <h3>{Name}</h3>
                <p>{formattedDate}</p>
            </div>
            <div className="event-meta">
                <span className={`status-icon ${status}`}></span>
            </div>
        </div>
    );
};

export default ProjectCard;