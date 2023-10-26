import {Component, OnInit, ViewChild} from '@angular/core';
import {OrderService} from "../../../shared/data/services/order.service";
import {Order} from "../../../../gen/ts/cheers/order/v1/order_service";
import {MatSort} from "@angular/material/sort";
import {Clipboard} from '@angular/cdk/clipboard';
import {PaymentService} from "../../../payments/data/payment.service";
import {RefundPaymentDialogComponent} from "../../ui/refund-payment-dialog/refund-payment-dialog.component";
import {MatTableDataSource} from "@angular/material/table";
import {MatDialog} from "@angular/material/dialog";
import {MatSnackBar} from "@angular/material/snack-bar";

@Component({
    selector: 'app-business-orders',
    templateUrl: './business-orders.component.html',
    styleUrls: ['./business-orders.component.sass']
})
export class BusinessOrdersComponent implements OnInit {
    displayedColumns: string[] = ['status', 'date', 'amount', 'firstName', 'lastName', 'email', 'more'];
    dataSource: MatTableDataSource<Order> = new MatTableDataSource();

    constructor(
        public dialog: MatDialog,
        private clipboard: Clipboard,
        private snackBar: MatSnackBar,
        private orderService: OrderService,
        private paymentService: PaymentService,
    ) {
    }

    @ViewChild(MatSort) sort: MatSort;

    ngAfterViewInit() {
        this.dataSource.sort = this.sort;
    }

    async ngOnInit() {
        const res = await this.orderService.getOrganizationOrders("")
        res.subscribe(data => {
            this.dataSource.data = data
        })
    }

    /** Copy order ID to clipboard */
    copyOrderID(orderID: string) {
        this.clipboard.copy(orderID);
        this.snackBar.open("Copied order id", "Hide", {
            duration: 3000
        })
    }

    /** Refund payment */
    refundPayment(paymentID: string) {
        const dialogRef = this.dialog.open(RefundPaymentDialogComponent, {
            panelClass: 'cheers-dialog'
        });

        dialogRef.afterClosed().subscribe(result => {
            if (result == true)
                this.paymentService.refundPayment(paymentID).subscribe(res => {
                    this.snackBar.open("Successful refund", "Hide", {
                        duration: 3000
                    })
                })
        });
    }

    /** Gets the total amount of all orders. */
    getTotalCost() {
        return this.dataSource.data.map(t => t.amount).reduce((acc, value) => acc + value, 0);
    }

    async onSearch(query: string) {
        const res = await this.orderService.getOrganizationOrders(query)
        res.subscribe(data => {
            this.dataSource.data = data
        })
    }
}