import React, { useState, useCallback } from "react";
import ReactDOM from 'react-dom';
import { render } from "react-dom";
import Gallery from "react-photo-gallery";
import Carousel, { Modal, ModalGateway } from "react-images";
import { photos } from "./photos";
import SimpleCard from "../Util/Card/Card"
function Home() {
  const [currentImage, setCurrentImage] = useState(0);
  const [viewerIsOpen, setViewerIsOpen] = useState(false);

  const openLightbox = useCallback((event, { photo, index }) => {
    setCurrentImage(index);
    setViewerIsOpen(true);
  }, []);

  const closeLightbox = () => {
    setCurrentImage(0);
    setViewerIsOpen(false);
  };

  return (
    <div>
      <div
        style={{
            position: 'fixed', left: '50%', top: '50%',
            transform: 'translate(-50%, -50%)', color: 'white', 
        }}
        >
          <SimpleCard />
      </div>
      <Gallery photos={photos} onClick={openLightbox} margin={2} targetRowHeight={500}/>
      <ModalGateway>
        {viewerIsOpen ? (
          <Modal onClose={closeLightbox}>
            <Carousel
              currentIndex={currentImage}
              views={photos.map(x => ({
                ...x,
                srcset: x.srcSet,
                caption: x.title
              }))}
            />
          </Modal>
        ) : null}
      </ModalGateway>
    </div>
  );
}
export default Home

ReactDOM.render(<Home />, document.getElementById('root'));
