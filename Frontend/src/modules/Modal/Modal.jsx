import React, {useState} from 'react';
import {Input, Modal, DatePicker, Select} from 'antd';

const MyModal = ({
                     open,
                     onOk,
                     title,
                     onCancel,
                     inputPlaceholder,
                     areaPlaceholder,
                     datePicker,
                     select
                 }) => {

    const {TextArea} = Input

    return (
        <Modal title={title} open={open} onOk={onOk} onCancel={onCancel}>
            <Input placeholder={inputPlaceholder} className='mb-3'/>
            <TextArea placeholder={areaPlaceholder}/>
            <div className='flex justify-between items-center mt-3'>
                {
                    datePicker ? <DatePicker placeholder='End date'/> : null
                }
                {
                    select ? <Select options={[
                        {value: 'open', label: <span>open</span>},
                        {value: 'on hold', label: <span>on hold</span>},
                        {value: 'in progress', label: <span>in progress</span>},
                        {value: 'done', label: <span>done</span>}
                    ]} className='w-40'/> : null
                }
            </div>

        </Modal>
    );
};

export default MyModal;