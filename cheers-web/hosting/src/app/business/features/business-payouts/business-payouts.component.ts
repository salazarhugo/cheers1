import { Component, OnInit } from '@angular/core';
import {AccountService} from "../../../shared/data/services/account.service";
import {Observable} from "rxjs";
import {Account} from "../../../../gen/ts/cheers/account/v1/account_service";

@Component({
  selector: 'app-business-payouts',
  templateUrl: './business-payouts.component.html',
  styleUrls: ['./business-payouts.component.sass']
})
export class BusinessPayoutsComponent implements OnInit {

  $account: Observable<Account | undefined>

  constructor(
      private accountService: AccountService,
  ) { }

  ngOnInit(): void {
    this.$account = this.accountService.getCurrentAccount()
  }

}
