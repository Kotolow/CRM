import React, { useState } from 'react';
import { GrNotification } from "react-icons/gr";
import { FiSearch, FiPlus } from 'react-icons/fi';
import MyModal from "../Modal/Modal";
import MyInput from "../MyInput/MyInput";
import { useParams } from "react-router-dom";

export default function Header({ title, buttonText, isSearched, isAdd, itemName, isDatePicker, isSelect, isPriority }) {
    const { id } = useParams(); 
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [itemTitle, setItemTitle] = useState(''); 
    const [itemDescription, setItemDescription] = useState('');
    const [taskStatus, setTaskStatus] = useState('open');
    const [dueDate, setDueDate] = useState(null);
    const [priority, setPriority] = useState('minor')

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleOk = async () => {
        // Create the task object
        const task = {
            title: itemTitle,
            description: itemDescription,
            status: taskStatus,
            due_date: dueDate ? dueDate.format('YYYY-MM-DD') : null,
            project_id: parseInt(id),
            priority: "minor",
            assigned_to: 1,
            comments: [],
            createdAt: new Date().toISOString(),
            updatedAt: new Date().toISOString(),
        };

        try {
            const response = await fetch('http://localhost/v1/projects/H%26h/tasks', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(task),
            });

            if (!response.ok) {
                const errorText = await response.text(); // Capture error text for debugging
                throw new Error(`HTTP error! status: ${response.status}, message: ${errorText}`);
            }

            const jsonResponse = await response.json();
            console.log("Task added successfully:", jsonResponse);
            // Reset form or show success message
            setItemTitle('');
            setItemDescription('');
            setTaskStatus('open');
            setDueDate(null);
            setIsModalOpen(false);
        } catch (error) {
            console.error("Error while adding task:", error.message); // Log the error message
        }
    };


    const handleCancel = () => {
        setIsModalOpen(false);
    };

    return (
        <div>
            <div className="top-bar">
                <a className="notifications flex items-center text-gray-700 hover:text-blue-500 p-2 bg-white hover:bg-gray-200 rounded-lg mr-4 cursor-pointer">
                    <GrNotification className="w-6 h-6" />
                </a>
                <a href='/profile' className="user-info flex items-center text-gray-700 hover:text-blue-500 p-2 bg-white hover:bg-gray-200 rounded-lg m-0 cursor-pointer">
                    <img src="profile_icon.png" alt="Profile" className="w-8 h-8 rounded-full mr-2" />
                    <span>Evan Yates</span>
                </a>
            </div>

            <div className="lower-bar flex justify-between items-center mb-6 mt-5">
                <div>
                    <h1 className="text-3xl font-bold text-gray-800">{title}</h1>
                </div>
                <div className="flex items-center space-x-4">
                    <div className="relative">
                        {isSearched && (
                            <div>
                                <input
                                    type="text"
                                    placeholder="Search"
                                    className="border rounded-lg p-2 pl-10 focus:outline-none focus:ring-2 focus:ring-blue-500"
                                />
                                <FiSearch className="absolute left-2 top-2 text-gray-400" />
                            </div>
                        )}
                    </div>
                    {isAdd && (
                        <div>
                            <button
                                className="bg-blue-500 text-white px-4 py-2 rounded-lg flex items-center hover:bg-blue-600"
                                onClick={showModal}
                            >
                                <FiPlus className="mr-2" /> {buttonText}
                            </button>
                        </div>
                    )}
                </div>
            </div>

            <MyModal
                open={isModalOpen}
                onOk={handleOk}
                onCancel={handleCancel}
                title="Add Item"
                inputPlaceholder={itemName}
                areaPlaceholder={itemDescription}
                datePicker={isDatePicker}
                select={isSelect}
                priority={isPriority}
                onInputChange={(value) => setItemTitle(value)}
                onTextAreaChange={(value) => setItemDescription(value)}
                onDateChange={(date) => setDueDate(date)}
                onSelectChange={(value) => setTaskStatus(value)}
                onSelectPriority={(value) => setPriority(value)}
            />
        </div>
    );
}
