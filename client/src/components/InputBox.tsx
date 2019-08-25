import * as React from 'react';
import { Input, Form, Select } from 'antd';
import brace from 'brace';
import AceEditor from 'react-ace';
import 'brace/mode/json';
import 'brace/theme/github';

const { TextArea } = Input;
const { Option } = Select;
 
export class InputBox extends React.Component {
    render() {
        return(
            <Form>
                <div className="select" style={{padding: '10px 0px 10px 0px'}}>
                    <Select
                        showSearch
                        style={{ width: 200 }}
                        placeholder="Select a person"
                        optionFilterProp="children"
                    >
                        <Option value="jack">Jack</Option>
                        <Option value="lucy">Lucy</Option>
                        <Option value="tom">Tom</Option>
                    </Select>
                </div>
                <div className="text" style={{padding: '10px 0px 10px 0px'}}>
                    <AceEditor
                        mode="json"
                        theme="github"
                        enableBasicAutocompletion={true}
                        height="250px"
                    />
                </div>
                
            </Form>
        );
    }
}