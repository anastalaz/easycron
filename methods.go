package easycron

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// defaultValues sets the easycron access token.
func (c *Client) defaultValues(v url.Values) url.Values {
	if v == nil {
		v = url.Values{}
	}
	v.Set("token", c.AccessToken)
	return v
}

// ListOutput holds the response from the list endpoint.
type ListOutput struct {
	Status   string `json:"status"`
	CronJobs []struct {
		CronJobID        string   `json:"cron_job_id"`
		CronJobName      string   `json:"cron_job_name"`
		Description      string   `json:"description"`
		UserID           string   `json:"user_id"`
		URL              string   `json:"url"`
		AuthUser         string   `json:"auth_user"`
		AuthPw           string   `json:"auth_pw"`
		CronExpression   string   `json:"cron_expression"`
		NumberFailedTime string   `json:"number_failed_time"`
		EpdsOccupied     string   `json:"epds_occupied"`
		EmailMe          string   `json:"email_me"`
		Sensitivity      string   `json:"sensitivity"`
		GroupID          string   `json:"group_id"`
		HTTPMethod       string   `json:"http_method"`
		HTTPHeaders      string   `json:"http_headers"`
		Posts            string   `json:"posts"`
		CustomTimeout    string   `json:"custom_timeout"`
		Criterion        string   `json:"criterion"`
		SuccessRegexp    string   `json:"success_regexp"`
		Wh               string   `json:"wh"`
		WhURL            string   `json:"wh_url"`
		WhHTTPMethod     string   `json:"wh_http_method"`
		WhData           []string `json:"wh_data"`
		Status           string   `json:"status"`
		Created          string   `json:"created"`
		Updated          string   `json:"updated"`
	} `json:"cron_jobs"`
}

// List gets all cron jobs of the account.
func (c *Client) List(v url.Values) (out *ListOutput, err error) {
	v = c.defaultValues(v)
	body, err := c.call("list", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// DetailOutput holds the response from the detail endpoint.
type DetailOutput struct {
	Status  string `json:"status"`
	CronJob struct {
		CronJobID        string   `json:"cron_job_id"`
		CronJobName      string   `json:"cron_job_name"`
		Description      string   `json:"description"`
		UserID           string   `json:"user_id"`
		URL              string   `json:"url"`
		AuthUser         string   `json:"auth_user"`
		AuthPw           string   `json:"auth_pw"`
		CronExpression   string   `json:"cron_expression"`
		NumberFailedTime string   `json:"number_failed_time"`
		EpdsOccupied     string   `json:"epds_occupied"`
		EmailMe          string   `json:"email_me"`
		Sensitivity      string   `json:"sensitivity"`
		GroupID          string   `json:"group_id"`
		HTTPMethod       string   `json:"http_method"`
		HTTPHeaders      string   `json:"http_headers"`
		Posts            string   `json:"posts"`
		CustomTimeout    string   `json:"custom_timeout"`
		Criterion        string   `json:"criterion"`
		SuccessRegexp    string   `json:"success_regexp"`
		Wh               string   `json:"wh"`
		WhURL            string   `json:"wh_url"`
		WhHTTPMethod     string   `json:"wh_http_method"`
		WhData           []string `json:"wh_data"`
		Status           string   `json:"status"`
		Created          string   `json:"created"`
		Updated          string   `json:"updated"`
	} `json:"cron_job"`
}

// Detail gets the details of a cron job.
func (c *Client) Detail(cronID int) (out *DetailOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))

	body, err := c.call("detail", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// AddOutput holds the response from the add endpoint.
type AddOutput struct {
	Status    string `json:"status"`
	CronJobID string `json:"cron_job_id"`
}

// Add creates a new cron job.
func (c *Client) Add(url string, cronExpression string, v url.Values) (out *AddOutput, err error) {
	v = c.defaultValues(v)
	v.Set("url", url)
	v.Set("cron_expression", cronExpression)

	body, err := c.call("add", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// TimezoneOutput holds the response from the timezone endpoint.
type TimezoneOutput struct {
	Status   string `json:"status"`
	Timezone string `json:"timezone"`
}

// Timezone gets the timezone of the account.
func (c *Client) Timezone() (out *TimezoneOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)

	body, err := c.call("timezone", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// EditOutput holds the response from the edit endpoint.
type EditOutput struct {
	Status    string `json:"status"`
	CronJobID string `json:"cron_job_id"`
}

// Edit gets the details of a cron job.
func (c *Client) Edit(cronID int, v url.Values) (out *EditOutput, err error) {
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))

	body, err := c.call("edit", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// EnableOutput holds the response from the enable endpoint.
type EnableOutput struct {
	Status    string `json:"status"`
	CronJobID string `json:"cron_job_id"`
}

// Enable enables a cron job.
func (c *Client) Enable(cronID int) (out *EnableOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))

	body, err := c.call("enable", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// DisableOutput holds the response from the disable endpoint.
type DisableOutput struct {
	Status    string `json:"status"`
	CronJobID string `json:"cron_job_id"`
}

// Disable disables a cron job.
func (c *Client) Disable(cronID int) (out *DisableOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))

	body, err := c.call("disable", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// DeleteOutput holds the response from the delete endpoint.
type DeleteOutput struct {
	Status    string `json:"status"`
	CronJobID string `json:"cron_job_id"`
}

// Delete deletes a cron job.
func (c *Client) Delete(cronID int) (out *DeleteOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))

	body, err := c.call("delete", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// LogsOutput holds the response from the logs endpoint.
type LogsOutput struct {
	Status string `json:"status"`
	Logs   []struct {
		CronJobID     string `json:"cron_job_id"`
		ScheduledTime string `json:"scheduled_time"`
		FireTime      string `json:"fire_time"`
		DoneTime      string `json:"done_time"`
		ExecuteTime   string `json:"execute_time"`
		HTTPCode      int    `json:"http_code"`
		Error         string `json:"error"`
		TotalTime     string `json:"total_time"`
		LogID         string `json:"log_id"`
	} `json:"logs"`
}

// Logs gets the last 10 logs of a cron job.
func (c *Client) Logs(cronID int) (out *LogsOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))

	body, err := c.call("logs", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// LogOutput holds the response from the log endpoint.
type LogOutput struct {
	Status          string `json:"status"`
	TruncatedOutput string `json:"truncated_output"`
}

// Log gets the truncated output of a specific log ID.
func (c *Client) Log(cronID int, logID string) (out *LogOutput, err error) {
	v := url.Values{}
	v = c.defaultValues(v)
	v.Set("id", strconv.Itoa(cronID))
	v.Set("log_id", logID)

	body, err := c.call("log", v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
