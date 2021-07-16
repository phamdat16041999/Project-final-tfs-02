<template>
  <div><div ref="paypal"></div></div>
</template>
<script>
export default {
  data() {
    return {
      order: {
        description: "Buy thing",
        amount: {
          currency_code: "USD",
          value: 1000
        }
      }
    };
  },
  mounted: function() {
    const script = document.createElement("script");
    const ClientID = "Aa6Rlkn9R-8IVOUQ2nog6MLZ-STb9kNB66o-6mEvk9IyrnXs7stvSudBOgwP0puc7Moa6gflZAmjSQU9";
    script.src = `https://www.paypal.com/sdk/js?client-id=${ClientID}`;
    script.addEventListener("load", this.setLoaded);
    document.body.appendChild(script);
  },
  methods: {
    setLoaded: function() {
      window.paypal
        .Buttons({
          createOrder: (data, actions) => {
            return actions.order.create({
              purchase_units: [this.order]
            });
          },
          onApprove: async (data, actions) => {
            const order = await actions.order.capture();
            alert(order)
            // ajax request
          },
          onError: err => {
            console.log(err);
          }
        })
        .render(this.$refs.paypal);
    }
  }
};
</script>
<style scoped>
</style>