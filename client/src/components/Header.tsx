import * as React from 'react';
import { Typography } from 'antd';

const { Text, Title } = Typography;

export interface HeaderProps {
    title: string;
}
 
export interface HeaderState {
    endpoint: string;
}
 
export class Header extends React.Component<HeaderProps, HeaderState> {
    state = {endpoint:  ""}
    render() { 
        return ( 
            <div>
                <Title underline>
                    easymock
                </Title>
                <Text code>
                    Usage area 
                </Text>
            </div>
         );
    }
}