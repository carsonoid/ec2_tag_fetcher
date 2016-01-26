# ec2_tag_fetcher

A simple golang utility to fetch the tags for the current instance and output them. Created specifically for salt-minion grains integration.

# Building

Simply clone the repo into your go environment and: 

`make all`

This will use docker to do the build.

# Running

It uses the [standard golang aws sdk method for credentials](https://github.com/aws/aws-sdk-go#configuring-credentials). The most direct way to run the tool is with environment variables:

`AWS_ACCESS_KEY_ID=AKID1234567890 AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY ./ec2_tag_fetcher`
