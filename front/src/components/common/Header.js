import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Typography from '@material-ui/core/Typography';

const Header = (props) => {
    return (
        <header>
            <nav className="nav-header">
                <h2>Hola!</h2>
                <img className="header-image" width="50px" height="50px" src="golang.png"></img>
            </nav>
        </header>
    )
}

export default Header;