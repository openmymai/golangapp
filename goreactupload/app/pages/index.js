import React, { useState, useRef } from "react"
import Image from "next/image"
import { Button, Container } from "react-bootstrap"
import axios from "axios"

const Index = () => {
  const [ image, setImage ] = useState(null)
  const [ loading, setLoading ] = useState(false)
  const [ uploadedImage, setUploadedImage ] = useState(null);
  const imageRef = useRef(null)

  const handleChangeImage = (e) => {
    setImage(e.target.files[0])
  }
  const handleUploadImage = async() => {
    try {
      setLoading(true)
      const formData = new FormData()
      formData.append("image", image)
      const res = await axios.post("http://localhost:8080", formData)
      if (res.data.data) {
        console.log(res.data)
        setUploadedImage(res.data.data)
      }
    } catch (error) {
      console.log(error)
    } finally {
      setImage(null)
      setLoading(false)
      document.getElementById("myUploadFile").value = null
    }
  }

  return (
    <>
      <br />
      <Container>
        <h4>Select Image</h4>
        <input 
          type="file"
          id="myUploadFile"
          ref={imageRef}
          onChange={handleChangeImage} 
        />
        <br />
        <br />
        
        {image && (
          <>
            <Image
              src={URL.createObjectURL(image)}
              width="300px"
              height="300px"
              alt="selected image..."
            />
            <br />
            <Button
              onClick={handleUploadImage}
            > 
            Upload
            </Button>
          </>
        )}
      </Container>
    </>
  )
}

export default Index