import React, { useState, useCallback } from "react";
import 'bootstrap/dist/css/bootstrap.min.css';
import Gallery from "react-photo-gallery";
import { photos } from "../../Home/photos";
import Carousel, { Modal, ModalGateway } from "react-images";
import Button from '@material-ui/core/Button';

class PrevForm extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      collageID: '',
      imageURL: '',
      imageName: '',
      imageDescription: '',
      collageSearchID: ''
    }

  }

  handleCollageIDChange = (e) => {
    this.setState({ collageID: e.target.value });
  };

  handleImageURLChange = (e) => {
    this.setState({ imageURL: e.target.value });
  };

  handleImageNameChange = (e) => {
    this.setState({ imageName: e.target.value });
  };

  handleImageDescriptionChange = (e) => {
    this.setState({ imageDescription: e.target.value });
  };

  handleCollageSearchChange = (e) => {
    this.setState({ collageSearchID: e.target.value });
  };

  onSubmit = (e) => {
    let payload = {
      CollageID: this.state.collageID,
      Name: this.state.imageName,
      Link: this.state.imageURL,
      Description: this.state.imageDescription
    }
    fetch('http://localhost:8080/apiv2/newphoto', {
    method: 'POST',
    body: JSON.stringify(payload)
  })
    return false
  }
  render() {
    return (

      <div className="card bg-light">

        <div className="card bg-light">
        <article className="card-body mx-auto" style={{ width: "500px" }}>
          <form>
            <h1 class="display-5">View Collage</h1>
            <div className="form-group">
              <label for="collageID">Enter the Collage ID</label>
              <input type="text" className="form-control" id="collageID" placeholder="Enter an ID" value={this.state.collageSearchID} onChange={this.handleCollageSearchChange}></input>
            </div>
            <Button variant="contained" color="secondary" href={"./display?id=" + this.state.collageSearchID}>View Collage </Button>
          </form>
        </article>
        </div>
        <Gallery photos={photos} margin={2} targetRowHeight={500} />

      </div>
    );
  }

}

export default PrevForm