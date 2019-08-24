import * as React from 'react';
import { Typography, Select } from 'antd';

const { Option } = Select;
const { Title } = Typography;
 
export class Content extends React.Component {
    render() { 
        return ( 
            <div>
                <Title level={2}>
                    Create an custom response endpoint!
                </Title>
            </div>
         );
    }
}
 
