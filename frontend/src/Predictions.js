import React, { useState, useEffect } from "react";
import { db } from "./firebase";
import { collection, getDocs } from "firebase/firestore";
import { Container, Typography, Card, CardContent } from "@mui/material";

const Predictions = () => {
  const [predictions, setPredictions] = useState([]);

  useEffect(() => {
    const fetchPredictions = async () => {
      const querySnapshot = await getDocs(collection(db, "predictions"));
      const data = querySnapshot.docs.map((doc) => ({
        id: doc.id,
        ...doc.data(),
      }));
      setPredictions(data);
    };

    fetchPredictions();
  }, []);

  return (
    <Container>
      <Typography variant="h4" style={{ margin: "20px 0" }}>
        MLB Predictions
      </Typography>
      {predictions.map((prediction) => (
        <Card key={prediction.id} style={{ marginBottom: "15px" }}>
          <CardContent>
            <Typography variant="body1">
              {JSON.stringify(prediction.prediction)}
            </Typography>
          </CardContent>
        </Card>
      ))}
    </Container>
  );
};

export default Predictions;
