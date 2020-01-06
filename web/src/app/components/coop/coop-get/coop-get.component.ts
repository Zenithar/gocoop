import { Component, OnInit } from '@angular/core';
import { CoopService } from '../../../services/coop.service';
import { NotificationsService } from 'angular2-notifications';

@Component({
  selector: 'app-coop-get',
  templateUrl: './coop-get.component.html',
  styleUrls: ['./coop-get.component.css']
})
export class CoopGetComponent implements OnInit {
  coopStatus: string = ""
  coop: any;

  constructor(
    private coopService: CoopService,
    private notificationService: NotificationsService
  ) { }

  ngOnInit() {
    this.getStatus();
    this.get();
  }

  getStatus(): void {
    this.coopService.getStatus().subscribe(
      (resp: string) => {
        this.coopStatus = resp;
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while getting the status', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

  get(): void {
    this.coopService.get().subscribe(
      (resp: any) => {
        this.coop = resp;
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while getting the coop', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

}
