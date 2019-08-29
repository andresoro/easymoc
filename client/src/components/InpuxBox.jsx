import React from 'react';
import { Button, Form, Select, Result } from 'antd';
import AceEditor from 'react-ace';
import 'brace/mode/json';
import 'brace/theme/github';
import axios from 'axios';

const Option = Select.Option;


export class InputBox extends React.Component {

    constructor() {
        super({});

        this.state = {
            endpoint: "",
            submitted: false,
            httpCode: 0,
            resp: false
        };
    }

    submit = () => {

        let text = this.refs.ace.editor.getValue();
        //make request
        let body = JSON.stringify(
            {code: this.state.httpCode, headers: {"Content-Type": "application/json"}, body: text}
        )
        
        axios.post('/gimme', body, {headers: {'Content-Type':'application/json'}})
            .then((response) => {
                this.setState({endpoint: response.data, submitted: true, resp: true})
            })
            .catch((error) => {
                this.setState({submitted: true, resp: false})
                console.log(error)
            })
    }

    optChange = (value) => {
        this.setState({httpCode: value})
    }

    clear = () => {
        this.setState({
            endpoint: "",
            submitted: false,
            httpCode: 0,
            body: "",
            resp: false,
        })
    }

    shouldComponentUpdate(nextProps, nextState) {
        if (this.state.body !== nextState.body) {
          return false
        } else {
          return true;
        }
      }

    render() {
        // if data has been sent to server, render the success/failure status
        if (this.state.submitted) {
            if (this.state.resp) {
                return (
                    <Result
                        key="pass"
                        status="success"
                        title="Endpoint successfully created"
                        extra={[
                        <Button key="pass" type="primary" onClick={this.clear}>
                            New Response
                        </Button>,
                        ]}
                    >
                    </Result>
                );
            } else {
                return (
                    <Result
                        key="fail"
                        status="error"
                        title="Submission Failed"
                        subTitle="Please try again in a few moments.."
                        extra={[
                        <Button key="fail" type="primary" onClick={this.clear}>
                            New Response
                        </Button>,
                        ]}
                    >   
                    </Result>
                );
            }
        // render the form if no submission has occurred yet or if form has been cleared    
        } else {
            return (
                <Form>
                <div className="select" style={{padding: '10px 0px 10px 0px'}}>
                    <Select
                        showSearch
                        style={{ width: 250 }}
                        placeholder="HTTP Status Code"
                        optionFilterProp="children"
                        onSelect={(value, event) => this.optChange(value)}
                    >
                        <Option value={100}>100: Continue</Option>
                        <Option value={101}>101: Switching Protocols</Option>
                        <Option value={102}>102: Processing</Option>
                        <Option value={200}>200: OK</Option>
                        <Option value={201}>201: Created</Option>
                        <Option value={202}>202: Accepted</Option>
                        <Option value={203}>203: Non-Authoritative Information</Option>
                        <Option value={204}>204: No Content</Option>
                        <Option value={205}>205: Reset Content</Option>
                        <Option value={206}>206: Partial Content</Option>
                        <Option value={207}>207: Multi-Status</Option>
                        <Option value={208}>208: Already Reported</Option>
                        <Option value={226}>226: IM Used</Option>
                        <Option value={300}>300: Multiple Choices</Option>
                        <Option value={301}>301: Moved Permanently</Option>
                        <Option value={302}>302: Found</Option>
                        <Option value={303}>303: See Other</Option>
                        <Option value={304}>304: Not Modified</Option>
                        <Option value={305}>305: Use Proxy</Option>
                        <Option value={306}>306: Switch Proxy</Option>
                        <Option value={307}>307: Temporary Redirect</Option>
                        <Option value={308}>308: Permanent Redirect</Option>
                        <Option value={400}>400: Bad Request</Option>
                        <Option value={401}>401: Unauthorized</Option>
                        <Option value={402}>402: Payment Required</Option>
                        <Option value={403}>403: Forbidden</Option>
                        <Option value={404}>404: Not Found</Option>
                        <Option value={405}>405: Method Not Allowed</Option>
                        <Option value={406}>406: Not Acceptable</Option>
                        <Option value={407}>407: Proxy Authentication Required</Option>
                        <Option value={408}>408: Request Timeout</Option>
                        <Option value={409}>409: Conflict</Option>
                        <Option value={410}>410: Gone</Option>
                        <Option value={411}>411: Length Required</Option>
                        <Option value={412}>412: Precondition Failed</Option>
                        <Option value={413}>413: Request Entity Too Large</Option>
                        <Option value={414}>414: Request-URI Too Long</Option>
                        <Option value={415}>415: Unsupported Media Type</Option>
                        <Option value={416}>416: Requested Range Not Satisfiable</Option>
                        <Option value={417}>417: Expectation Failed</Option>
                        <Option value={418}>418: I'm a teapot</Option>
                        <Option value={420}>420: Enhance Your Calm</Option>
                        <Option value={422}>422: Unprocessable Entity</Option>
                        <Option value={423}>423: Locked</Option>
                        <Option value={424}>424: Failed Dependency</Option>
                        <Option value={425}>425: Unordered Collection</Option>
                        <Option value={426}>426: Upgrade Required</Option>
                        <Option value={428}>428: Precondition Required</Option>
                        <Option value={429}>429: Too Many Requests</Option>
                        <Option value={431}>431: Request Header Fields Too Large</Option>
                        <Option value={444}>444: No Response</Option>
                        <Option value={449}>449: Retry With</Option>
                        <Option value={450}>450: Blocked by Windows Parental Controls</Option>
                        <Option value={499}>499: Client Closed Request</Option>
                        <Option value={500}>500: Internal Server Error</Option>
                        <Option value={501}>501: Not Implemented</Option>
                        <Option value={502}>502: Bad Gateway</Option>
                        <Option value={503}>503: Service Unavailable</Option>
                        <Option value={504}>504: Gateway Timeout</Option>
                        <Option value={505}>505: HTTP Version Not Supported</Option>
                        <Option value={506}>506: Variant Also Negotiates</Option>
                        <Option value={507}>507: Insufficient Storage</Option>
                        <Option value={509}>509: Bandwidth Limit Exceeded</Option>
                        <Option value={510}>510: Not Extended</Option>
                    </Select>
                </div>
                <div className="text" style={{padding: '10px 0px 10px 0px'}}>
                    <AceEditor
                        mode="json"
                        theme="github"
                        enableBasicAutocompletion={true}
                        height="250px"
                        ref="ace"
                    />
                </div>
                <div className="button">
                    <Button type="primary" onClick={this.submit}>Submit</Button>
                </div>
                
            </Form>
            );
        }
    }
}