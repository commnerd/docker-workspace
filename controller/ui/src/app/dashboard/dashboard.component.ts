import { FormBuilder } from '@angular/forms'
import { Component, OnInit } from '@angular/core'

import { Proxy } from '../models/proxy'

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  dashboardForm = this._fb.group({
    proxies: this._fb.array([])
  })

  constructor(private _fb: FormBuilder) {}

  ngOnInit() {
  	this.addProxy(new Proxy)
  }

  addProxy(proxy: Proxy) {
    this.dashboardForm.get("proxies").push(this._fb.group({
      port: [proxy.port],
      env_base_path: [proxy.env_base_path],
      site_base_path: [proxy.site_base_path]
    }));
  }
}
