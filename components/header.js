import React, { Component} from 'react';
import { Toolbar, ToolbarGroup, ToolbarSeparator, ToolbarTitle } from 'material-ui/Toolbar';

export default class extends Component {
    render() {
        return (
            <Toolbar className="appBar">
                <style jsx global>{`
                    .appBar {
                        padding: 50px 5px 5px;
                    }

                    .appBar span.brand {
                        color: green;
                    }
                `}</style>
                <ToolbarGroup firstChild={true}>
                    <ToolbarTitle className="brand" text="Hithes Direct" />
                </ToolbarGroup>
                <ToolbarGroup>
                    <ToolbarSeparator />
                </ToolbarGroup>
            </Toolbar>
        );
    }
}