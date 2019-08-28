import * as React from 'react';
import { Typography } from 'antd';

const { Title } = Typography;

export class Header extends React.Component {
    render() {
        return(
           <Title level={2} underline>easymock</Title>
        );
    }
}