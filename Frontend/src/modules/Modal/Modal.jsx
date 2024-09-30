import React from 'react';
import { Input, Modal, DatePicker, Select } from 'antd';

const MyModal = ({
                     open,
                     onOk,
                     title,
                     onCancel,
                     inputPlaceholder,
                     areaPlaceholder,
                     datePicker,
                     select,
                     onInputChange,
                     onTextAreaChange,
                     onDateChange,
                     onSelectChange,
                 }) => {
    const { TextArea } = Input;

    return (
        <Modal title={title} open={open} onOk={onOk} onCancel={onCancel}>
            <Input
                placeholder={inputPlaceholder}
                className='mb-3'
                onChange={(e) => onInputChange(e.target.value)} // Handle title input change
            />
            <TextArea
                placeholder={areaPlaceholder}
                onChange={(e) => onTextAreaChange(e.target.value)} // Handle description input change
            />
            <div className='flex justify-between items-center mt-3'>
                {datePicker && (
                    <DatePicker
                        placeholder='End date'
                        onChange={onDateChange} // Handle date change
                    />
                )}
                {select && (
                    <Select
                        options={[
                            { value: 'open', label: <span>open</span> },
                            { value: 'on hold', label: <span>on hold</span> },
                            { value: 'in progress', label: <span>in progress</span> },
                            { value: 'done', label: <span>done</span> },
                        ]}
                        className='w-40'
                        onChange={onSelectChange} // Handle select change
                    />
                )}
            </div>
        </Modal>
    );
};

export default MyModal;
