// Copyright (c) 2020 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: Sep 2020
// This file contains the JS functions for index.html

'use-strict'

/*
* This function calculates the area of a trapezoid
*/
function calculate () {
  // debugging statements
  console.log('button clicked')

  // grab input from the text box
  const aLength = parseFloat(document.getElementById('a-length').value)
  const bLength = parseFloat(document.getElementById('b-length').value)
  const height = parseFloat(document.getElementById('height').value)

  // calculate the Volume
  const volume = (aLength * bLength) * height / 3

  // Output volume to the user
  document.getElementById('answer').innerHTML = "The volume is: " + volume + " cm³"
}
