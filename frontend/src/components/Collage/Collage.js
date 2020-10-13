import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import GridList from '@material-ui/core/GridList';
import GridListTile from '@material-ui/core/GridListTile';
import GridListTileBar from '@material-ui/core/GridListTileBar';
import ListSubheader from '@material-ui/core/ListSubheader';
import IconButton from '@material-ui/core/IconButton';
import InfoIcon from '@material-ui/icons/Info';
import tileData from './DemoData';
import 'bootstrap/dist/css/bootstrap.min.css';

class CollageGrid extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      data: []
    }

  }

  componentDidMount() {
    const urlParams = new URLSearchParams(window.location.search);
    const collageID = urlParams.get('id');
    fetch(`http://localhost:8080/apiv2/collage/${collageID}`)
        .then(res => res.json())
        .then(result => {
          this.setState({ data: result })
        })
  }

  render() {
    if (!this.state.data) {
      return <h1>No collage found :(</h1>
    }
    return (
      <div style={{

        display: 'flex',
        flexWrap: 'wrap',
        justifyContent: 'space-around',
        overflow: 'hidden',
        backgroundColor: "theme.palette.background.paper",

      }}>
        <GridList cellHeight={300} spacing={10} style={{
              width: 1000,
              height: 1000,
        }}>
          <GridListTile key="Subheader" cols={4} style={{ height: 'auto' }}>
            <ListSubheader component="div"></ListSubheader>
          </GridListTile>
          {this.state.data.map((tile) => (
            <GridListTile key={tile.img}>
              <img src={tile.link} alt={tile.name} />
              <GridListTileBar
                title={tile.name}
                subtitle={tile.description}
              />
            </GridListTile>
          ))}
        </GridList>
      </div>
    );
  }
}

export default CollageGrid;