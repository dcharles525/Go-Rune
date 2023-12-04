import Controller from '@ember/controller';

export default class ApplicationController extends Controller {

  testData = {
    labels: ['test', 'test' ,'test', 'test', 'test'],
    datasets: [{
      label: 'My First Dataset',
      data: [1,4,10,6,7],
      borderColor: [
        'rgba(97, 97, 97, 1)'
      ],
      backgroundColor: [
        'rgba(97, 97, 97, 0.2)'
      ]
    }],
  };

  testOptions = {
    responsive: true,
    legend: {
      labels: {
        fontSize: 30,
          fontColor: 'rgb(248, 248, 248)'
      }
    },
    scales: {
      yAxes: [{
        ticks: {
          fontSize: 20,
          fontColor: 'rgb(248, 248, 248)'
        }
      }],
      xAxes: [{
        ticks: {
          fontSize: 20,
          fontColor: 'rgb(248, 248, 248)'
        }
      }]
    }
  }

}
