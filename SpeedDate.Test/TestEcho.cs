﻿using System;
using System.Collections.Generic;
using System.Net;
using System.Text;
using System.Threading;
using NUnit.Framework;
using Shouldly;
using SpeedDate.Client;
using SpeedDate.ClientPlugins.Peer.Echo;
using SpeedDate.Configuration;

namespace SpeedDate.Test
{
    [TestFixture]
    public class TestEcho
    {
        [Test]
        public void TestSingleEcho()
        {
            const string message = "MyTestMessage12345";

            var doneEvent = new AutoResetEvent(false);

            var client = new SpeedDateClient();
            client.Started += () =>
            {
                client.IsConnected.ShouldBeTrue();

                client.GetPlugin<EchoPlugin>().Send(message,
                    echo =>
                    {
                        echo.ShouldBe(message);
                        doneEvent.Set();
                    },
                    error =>
                    {
                        Should.NotThrow(() => throw new Exception(error));
                    });
            };

            client.Start(new DefaultConfigProvider(
                new NetworkConfig(IPAddress.Loopback, SetUp.Port), //Connect to port
                new PluginsConfig("SpeedDate.ClientPlugins.Peer*"))); //Load peer-plugins only

            doneEvent.WaitOne(TimeSpan.FromSeconds(30)).ShouldBeTrue(); //Should be signaled
        }

        [Test]
        public void TestMultiClientEcho()
        {
            var numberOfClients = 100;


            var doneEvent = new AutoResetEvent(false);

            for (var clientNumber = 0; clientNumber < numberOfClients; clientNumber++)
            {
                ThreadPool.QueueUserWorkItem(state =>
                {
                    var client = new SpeedDateClient();
                    client.Started += () =>
                    {
                        client.IsConnected.ShouldBeTrue();

                        client.GetPlugin<EchoPlugin>().Send("Hello from " + state,
                            echo =>
                            {
                                echo.ShouldBe("Hello from " + state);

                                if (Interlocked.Decrement(ref numberOfClients) == 0)
                                    doneEvent.Set();

                            },
                            error =>
                            {
                                Should.NotThrow(() => throw new Exception(error));
                            });
                    };

                    client.Start(new DefaultConfigProvider(
                        new NetworkConfig(IPAddress.Loopback, SetUp.Port), //Connect to port
                        new PluginsConfig("SpeedDate.ClientPlugins.Peer*"))); //Load peer-plugins only

                }, clientNumber);
            }


            doneEvent.WaitOne(TimeSpan.FromSeconds(30)).ShouldBeTrue(); //Should be signaled
        }
    }
}